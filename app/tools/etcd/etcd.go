package etcd

import (
	"sync"

	"github.com/stone2401/light-gateway/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var etcdClient *clientv3.Client
var etcdOnce sync.Once

func GetEtcdClient() *clientv3.Client {
	etcdOnce.Do(func() {
		var err error
		etcdClient, err = clientv3.New(clientv3.Config{Endpoints: config.Config.Endpoints})
		if err != nil {
			panic(err)
		}
	})
	return etcdClient
}
