package proxy

import (
	"net/http"
	"sync"

	"github.com/stone2401/light-gateway-kernel/pcore"
)

type Status struct {
	online bool
	rw     sync.RWMutex
}

func NewStatus() *Status {
	return &Status{
		online: false,
		rw:     sync.RWMutex{},
	}
}

func (s *Status) StatusHeadler(ctx *pcore.Context) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	if s.online {
		ctx.Next()
		return
	}
	http.NotFound(ctx.Response, ctx.Request)
	ctx.Abort()
}

// 开启
func (s *Status) Open() {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.online = true
}

// 关闭
func (s *Status) Close() {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.online = false
}
