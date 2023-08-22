package public

import (
	"errors"
	"math/rand"
	"sync"
)

type RandomBalance struct {
	length int
	rss    []string
	mx     sync.RWMutex
}

func (r *RandomBalance) Get() (string, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()
	if len(r.rss) == 0 {
		return "", errors.New("rss is empty")
	}
	return r.rss[rand.Intn(r.length)], nil
}

func (r *RandomBalance) Add(params ...string) error {
	if r.rss == nil {
		r.rss = make([]string, 0)
	}
	if len(params) == 0 {
		return errors.New("params is empty")
	}
	addr := params[0]
	r.mx.Lock()
	defer r.mx.Unlock()
	r.rss = append(r.rss, addr)
	r.length++
	return nil
}

func (r *RandomBalance) Delete(params ...string) {
	for i := range r.rss {
		if r.rss[i] == params[0] {
			r.mx.Lock()
			r.rss = append(r.rss[:i], r.rss[i+1:]...)
			r.mx.Unlock()
		}
	}
}
