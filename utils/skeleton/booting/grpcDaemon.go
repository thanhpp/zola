package booting

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/pinezapple/LibraryProject20201/skeleton/configs"
	"github.com/pinezapple/LibraryProject20201/skeleton/libs"
	"github.com/pinezapple/LibraryProject20201/skeleton/libs/etcd"
	"github.com/pinezapple/LibraryProject20201/skeleton/logger"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"

	etcdClient "go.etcd.io/etcd/clientv3"
	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials"
)

// GRPCService daemon template for gRPC service
func GRPCService(parentCtx context.Context,
	serviceName string,
	client *etcdClient.Client,
	conf configs.GRPCServerConfig,
	register func(*grpc.Server),
	opts ...grpc.ServerOption,
) (daemon model.Daemon, err error) {

	lg := logger.MustGet(serviceName + "grpc Server")

	// try to listening to bind address
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		return
	}
	s := grpc.NewServer(opts...)
	register(s)

	go func() {
		if e := s.Serve(listen); e != nil {
			logger.LogErr(lg, e)
			return
		}
	}()

	daemon = func() {
		//	defer lg.Warn("Gracefully stop gRPC Server")

		//	lg.Infof("gRPC listening(port: %d)", conf.Port)

		logger.LogInfo(lg, "gRPC start listening at port"+strconv.Itoa(conf.Port))

		// wait a while for server booting
		time.Sleep(400 * time.Millisecond)

		// keep alive with etcd
		go etcd.KeepAliveService(parentCtx, client, serviceName, conf.PublicIP, conf.Port)

		// wait ctx out
		<-parentCtx.Done()

		// Stop server now
		s.Stop()
	}
	return
}

// loadRPCServerOptions loading rpc server opts
func loadRPCServerOptions(crt, key, clientCA string) ([]grpc.ServerOption, error) {
	// load certificates
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, err
	}

	clientCAs, err := libs.LoadCACertPool([]string{clientCA})
	if err != nil {
		return nil, err
	}

	// server options
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewTLS(&tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{certificate},
			ClientCAs:    clientCAs,
		})),
	}

	return opts, nil
}
