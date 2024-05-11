package etcd

import (
	"sync"

	"github.com/stone2401/light-gateway-kernel/pcore"
)

var monitor pcore.Monitor
var monce sync.Once

func GetMonitor() pcore.Monitor {
	monce.Do(func() {
		monitor = pcore.NewEtcdMonitor(GetEtcdClient())
	})
	return monitor
}
