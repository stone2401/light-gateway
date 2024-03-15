package public

import (
	"context"
	"log"
	"sync"
)

const (
	// 随机轮询
	RANDOMBALANCE = iota
	// 顺序轮询
	ROUNDROBINBALANCE
	// 权重轮询
	WEIGHTROUNDBALANCE
)

type Balance interface {
	Get(string) (string, error)
	Add(params ...string) error
	Delete(params ...string)
}

func NewBalance(bal int, suffix string, ctx context.Context) Balance {
	var balance Balance
	switch bal {
	case 0:
		balance = &RandomBalance{
			length: 0,
			rss:    make([]string, 0),
			mx:     sync.RWMutex{},
		}
	case 1:
		balance = &RoundRobinBalance{
			curIndex: 0,
			rss:      make([]string, 0),
			mx:       sync.RWMutex{},
		}
	case 2:
		balance = &WeightRoundBalance{
			curIndex: 0,
			rss:      make([]*WeightNode, 0),
			mx:       sync.RWMutex{},
		}
	default:
		balance = &RandomBalance{
			length: 0,
			rss:    make([]string, 0),
			mx:     sync.RWMutex{},
		}
	}
	go Watch(ctx, balance, suffix)
	return balance
}

func Watch(ctx context.Context, balance Balance, suffix string) error {
	channels := PREFIX + GetRedis().SuffixChack(suffix)
	keys := GetRedis().GetKeys(GetRedis().SuffixChack(suffix))
	for _, v := range keys {
		balance.Add(v, GetRedis().GetValue(v))
	}
	pub := GetRedis().PSubscribe(channels)
	defer pub.Close()
	for {
		log.Println(balance.Get("hello"))
		select {
		case <-ctx.Done():
			return nil
		case msg := <-pub.Channel():
			switch msg.Payload {
			case "set":
				log.Println(msg.Pattern)
				balance.Add(msg.Channel, GetRedis().GetValue(msg.Channel))
			case "expired":
				balance.Delete(msg.Channel)
				log.Println(msg.Pattern)
			case "del":
				balance.Delete(msg.Channel)
				log.Println(msg.Pattern)
			default:
				log.Println("收到了，但是没有触发", msg.Payload)
			}
		}
	}
}
