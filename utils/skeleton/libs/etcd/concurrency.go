package etcd

import (
	"context"

	"github.com/pinezapple/LibraryProject20201/skeleton/logger"
	"github.com/pinezapple/LibraryProject20201/skeleton/model"

	etcdClient "go.etcd.io/etcd/clientv3"
	etcdSync "go.etcd.io/etcd/clientv3/concurrency"
)

// etcd's default of 60s is too long, shortened to 5s
const defaultSessionTTL = 5

// WatchSession is a watchdog that keeps track of your session status
// giving you tools to ensure fault-tolerant when using etcd mutex
func WatchSession(
	ctx context.Context,
	client *etcdClient.Client,
	serviceName string,
	setupLocks func(*etcdSync.Session,
	) (err error),
	opts ...SessionOption) (daemon model.Daemon, err error) {
	lg := logger.MustGet(serviceName)

	ops := &sessionOptions{
		leaseID: etcdClient.NoLease, // default to create a new lease
		ttl:     defaultSessionTTL,
		ctx:     client.Ctx(), // use client's context instead of app contextt
	}
	for _, opt := range opts {
		opt(ops)
	}

	// setup new lease and keepalive on that lease
	var session *etcdSync.Session
	session, err = etcdSync.NewSession(client,
		etcdSync.WithContext(ops.ctx),
		etcdSync.WithTTL(ops.ttl),
		etcdSync.WithLease(ops.leaseID))
	if err != nil {
		return
	}

	if err = setupLocks(session); err != nil {
		return
	}

	daemon = func() {
		select {
		case <-ctx.Done(): // graceful shutdown
			return
		case _, open := <-session.Done():
			if !open { // fatally stop the process
				if ops.shouldPanic {
					panic("Fatal: ETCD concurrency session is closed")
				} else {
					//	lg.Error("Fatal: ETCD concurrency session is closed")
					logger.LogErr(lg, err)
				}
			}

		}
	}
	return
}

type sessionOptions struct {
	ttl         int
	leaseID     etcdClient.LeaseID
	ctx         context.Context
	shouldPanic bool
}

// SessionOption configures Session.
type SessionOption func(*sessionOptions)

// WithTTL configures the session's TTL in seconds.
// If TTL is <= 0, the default 60 seconds TTL will be used.
func WithTTL(ttl int) SessionOption {
	return func(so *sessionOptions) {
		if ttl > 0 {
			so.ttl = ttl
		}
	}
}

// WithLease specifies the existing leaseID to be used for the session.
// This is useful in process restart scenario, for example, to reclaim
// leadership from an election prior to restart.
func WithLease(leaseID etcdClient.LeaseID) SessionOption {
	return func(so *sessionOptions) {
		so.leaseID = leaseID
	}
}

func WithPanic() SessionOption {
	return func(so *sessionOptions) {
		so.shouldPanic = true
	}

}
