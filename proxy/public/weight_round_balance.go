package public

import (
	"errors"
	"strconv"
	"sync"
)

type WeightRoundBalance struct {
	curIndex int
	rss      []*WeightNode
	mx       sync.RWMutex
}

type WeightNode struct {
	addr            string
	weight          int // 权重值
	currentWeight   int // 当前权重
	effectiveWeight int // 有效权重
}

func (w *WeightRoundBalance) Get() (string, error) {
	total := 0
	var best *WeightNode
	for _, value := range w.rss {
		total += value.effectiveWeight
		value.currentWeight += value.effectiveWeight
		if best == nil || value.currentWeight > best.currentWeight {
			best = value
		}
	}
	if best == nil {
		return "", errors.New("[!] error")
	}
	best.currentWeight -= total
	return best.addr, nil
}

func (w *WeightRoundBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params len need 2")
	}
	i, err := strconv.Atoi(params[1])
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: i, effectiveWeight: i}
	w.rss = append(w.rss, node)
	return nil
}

func (r *WeightRoundBalance) Delete(params ...string) {
	for i := range r.rss {
		if r.rss[i].addr == params[0] {
			r.mx.Lock()
			r.rss = append(r.rss[:i], r.rss[i+1:]...)
			r.mx.Unlock()
		}
	}
}
