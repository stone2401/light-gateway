package public

import (
	"errors"
	"sync"
)

type RoundRobinBalance struct {
	curIndex int
	rss      []string
	mx       sync.RWMutex
}

func (r *RoundRobinBalance) Get() (string, error) {
	if len(r.rss) == 0 {
		return "", errors.New("rss is empty")
	}
	if r.curIndex >= len(r.rss) {
		r.curIndex = 0
	}
	curAddir := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1) % len(r.rss)
	return curAddir, nil
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if r.rss == nil {
		r.rss = make([]string, 0)
	}
	if len(params) == 0 {
		return errors.New("params is empty")
	}
	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RoundRobinBalance) Delete(params ...string) {
	for i := range r.rss {
		if r.rss[i] == params[0] {
			r.mx.Lock()
			r.rss = append(r.rss[:i], r.rss[i+1:]...)
			r.mx.Unlock()
		}
	}
}
