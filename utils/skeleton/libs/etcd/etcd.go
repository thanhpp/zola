package etcd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/pinezapple/LibraryProject20201/skeleton/libs"
	"github.com/pinezapple/LibraryProject20201/skeleton/logger"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"

	etcdClient "go.etcd.io/etcd/clientv3"
	etcdNaming "go.etcd.io/etcd/clientv3/naming"
	naming "google.golang.org/grpc/naming"
)

const DefaultTimeout = 10 * time.Second

// LoadEtcdClient load etcd client from scratch
func LoadEtcdClient(etcdEndpoints, etcdCertFile, etcdKeyFile, etcdCaFile string) (client *etcdClient.Client, err error) {
	var tlsConfig *tls.Config

	if etcdCertFile != "" && etcdKeyFile != "" && etcdCaFile != "" {
		// Load client cert
		cert, err := tls.LoadX509KeyPair(etcdCertFile, etcdKeyFile)
		if err != nil {
			return nil, err
		}

		caCertPool, err := libs.LoadCACertPool([]string{etcdCaFile})
		if err != nil {
			return nil, err
		}

		// Setup tls config transport
		tlsConfig = &tls.Config{
			Certificates:             []tls.Certificate{cert},
			RootCAs:                  caCertPool,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}
		tlsConfig.BuildNameToCertificate()
	}

	// setup etcd config
	return etcdClient.New(etcdClient.Config{
		Endpoints:   strings.Split(etcdEndpoints, ","),
		TLS:         tlsConfig,
		DialTimeout: DefaultTimeout,
	})
}

// KeepAliveService makes service keep alive to etcd (service discovery).
func KeepAliveService(parentCtx context.Context,
	client *etcdClient.Client,
	serviceName string,
	publicHost string,
	publicPort int,
) {
	lg := logger.MustGet(serviceName + "keep alive")

	for {
		select {
		case <-parentCtx.Done():
			return

		case <-time.After(100 * time.Millisecond):
		}

		lease, err := client.Grant(parentCtx, 2) // time to live: 2 seconds
		if err != nil || lease == nil {
			//			lg.Error(err)
			logger.LogErr(lg, err)
			continue
		}

		// attach key with lease
		resolver := &etcdNaming.GRPCResolver{Client: client}
		err = resolver.Update(parentCtx, serviceName, naming.Update{ // nolint: staticcheck
			Op:   naming.Add,
			Addr: fmt.Sprintf("%s:%d", publicHost, publicPort),
		}, etcdClient.WithLease(lease.ID))
		if err != nil {
			//			lg.Error(err)
			logger.LogErr(lg, err)
			continue
		}

		// keep alive lease
		keepAliveCh, err := client.KeepAlive(parentCtx, lease.ID)
		if err != nil {
			//			lg.Error(err)
			logger.LogErr(lg, err)
			continue
		}

	L:
		for {
			select {
			case rp, opened := <-keepAliveCh:
				// The returned "LeaseKeepAliveResponse" channel closes if underlying keep
				// alive stream is interrupted in some way the client cannot handle itself;
				// given context "ctx" is canceled or timed out. "LeaseKeepAliveResponse"
				// from this closed channel is nil.
				if rp == nil || !opened {
					break L
				}

			case <-parentCtx.Done():
				return
			}
		}
	}
}

// Watch for value change and notify.
func Watch(ctx context.Context, client *etcdClient.Client, key string, callback func() error) (daemon model.Daemon, err error) {
	lg := logger.MustGet("ConfigWatcher" + key)

	daemon = func() {
		watchan := client.Watch(ctx, key)
		for {
			select {

			case v := <-watchan:
				if v.Err() != nil {
					continue
				}

				// handle update conf
				if e := callback(); e != nil {
					//				lg.Error(e)
					logger.LogErr(lg, err)
				}

			case <-ctx.Done():
				return
			}
		}
	}
	return
}

// Get value and notify.
func Get(client *etcdClient.Client, key string, expect interface{}, callback func(interface{}) error) (err error) {
	var resp *etcdClient.GetResponse
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	resp, err = client.Get(ctx, key)
	cancel()
	if err == nil {
		if resp.Count <= 0 {
			err = fmt.Errorf("Cannot find key %s in etcd", key)
			return
		}

		if err = json.Unmarshal(resp.Kvs[0].Value, expect); err == nil {
			err = callback(expect)
			return
		}
	}

	return
}
