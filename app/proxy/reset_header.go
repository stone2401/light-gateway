package proxy

import (
	"strings"
	"sync"

	"github.com/stone2401/light-gateway-kernel/pcore"
)

type ResetHeader struct {
	rw      sync.RWMutex
	editMap map[string]string
	addMap  map[string]string
	delMap  map[string]string
}

func NewResetHeader() *ResetHeader {
	return &ResetHeader{
		rw:      sync.RWMutex{},
		editMap: map[string]string{},
		addMap:  map[string]string{},
		delMap:  map[string]string{},
	}
}

// headerTransfor 结构 add header value
func (r *ResetHeader) Set(headerTransfor string) {
	r.rw.Lock()
	defer r.rw.Unlock()
	headers := strings.Split(headerTransfor, "\n")
	for _, header := range headers {
		if len(strings.Split(header, " ")) != 3 {
			continue
		}
		switch strings.Split(header, " ")[0] {
		case "add":
			r.addMap[strings.Split(header, " ")[1]] = strings.Split(header, " ")[2]
		case "edit":
			r.editMap[strings.Split(header, " ")[1]] = strings.Split(header, " ")[2]
		case "del":
			r.delMap[strings.Split(header, " ")[1]] = strings.Split(header, " ")[2]
		}
	}
}

func (r *ResetHeader) ResetHeader(ctx *pcore.Context) {
	r.rw.RLock()
	defer r.rw.RUnlock()
	for key, value := range r.addMap {
		ctx.Request.Header.Add(key, value)
	}
	for key, value := range r.editMap {
		ctx.Request.Header.Set(key, value)
	}
	for key := range r.delMap {
		ctx.Request.Header.Del(key)
	}
	ctx.Next()
}

func (r *ResetHeader) Reset(ctx *pcore.Context) {
	r.rw.Lock()
	defer r.rw.Unlock()
	r.addMap = map[string]string{}
	r.editMap = map[string]string{}
	r.delMap = map[string]string{}
}
