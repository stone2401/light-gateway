package public

import (
	"errors"
	"sync"

	"github.com/golang/groupcache/consistenthash"
)

type HashBalance struct {
	hash *consistenthash.Map
	mu   *sync.RWMutex
	rss  []string
}

func NewHashBalance() *HashBalance {
	return &HashBalance{
		hash: consistenthash.New(0, nil),
		mu:   &sync.RWMutex{},
		rss:  make([]string, 0),
	}
}

func (h *HashBalance) Get(hash string) (string, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.rss) == 0 {
		return "", errors.New("rss is empty")
	}
	return h.hash.Get("asklfhaslksdjasp;ourqpwoalskxc,mnlakhd"), nil
}

func (h *HashBalance) Add(params ...string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(params) == 0 {
		return errors.New("params is empty")
	}
	addr := params[0]
	h.rss = append(h.rss, addr)
	h.hash = consistenthash.New(len(h.rss), nil)
	for _, v := range h.rss {
		h.hash.Add(v)
	}
	return nil
}

func (h *HashBalance) Delete(params ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, item := range h.rss {
		if item != params[0] {
			h.rss = append(h.rss, item)
		}
	}
	h.hash = consistenthash.New(len(h.rss), nil)
	for _, v := range h.rss {
		h.hash.Add(v)
	}
}
