package proxy

import (
	"strings"
	"sync"

	"github.com/stone2401/light-gateway-kernel/pcore"
)

type AccessControl struct {
	rw        sync.RWMutex
	WhiteList map[string]struct{}
	BlackList map[string]struct{}
}

func NewAccessControl() *AccessControl {
	return &AccessControl{
		WhiteList: map[string]struct{}{},
		BlackList: map[string]struct{}{},
		rw:        sync.RWMutex{},
	}
}

func (a *AccessControl) AddWhiteList(whiteList string) {
	a.rw.Lock()
	defer a.rw.Unlock()
	whites := strings.Split(whiteList, ",")
	for _, white := range whites {
		a.WhiteList[white] = struct{}{}
	}
}

func (a *AccessControl) AddBlackList(blackList string) {
	a.rw.Lock()
	defer a.rw.Unlock()
	blacks := strings.Split(blackList, ",")
	for _, black := range blacks {
		if black == "" {
			continue
		}
		a.BlackList[black] = struct{}{}
	}
}

func (a *AccessControl) IsAllow(ip string) bool {
	a.rw.RLock()
	defer a.rw.RUnlock()
	if _, ok := a.WhiteList[ip]; ok {
		return true
	}
	if _, ok := a.BlackList[ip]; ok {
		return false
	}
	return true
}

func (a *AccessControl) AccessControlHeadler(ctx *pcore.Context) {
	if !a.IsAllow(ctx.Request.RemoteAddr) {
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (a *AccessControl) Reset() {
	a.rw.Lock()
	defer a.rw.Unlock()
	a.WhiteList = map[string]struct{}{}
	a.BlackList = map[string]struct{}{}
}
