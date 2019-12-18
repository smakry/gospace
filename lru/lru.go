// LRU caching scheme
// queue storage cache
// map storage has cache

package lru

import (
	"gospace/queue"
//	"fmt"
)

const lruMaxSize = 5

type PageUse struct {
	Addr	map[interface{}]*queue.Node
	UseQu	*queue.Queue
}

func NewPageUse() *PageUse {
	return &PageUse{
		Addr:	make(map[interface{}]*queue.Node),
		UseQu:	queue.NewQueue(),
	}
}

func (p *PageUse) MaxSize() int {
	return lruMaxSize
}

func (p *PageUse) View(id interface{}) {
	v, ok := p.Addr[id]
	p.Addr[id] = (p.UseQu.Push(id)).(*queue.Node)

	if ok {
		//del address change, update next node address
		if del, ok := p.UseQu.DelNode(v); ok && del != nil {
			p.Addr[del.Data] = del
		}
	} else {
		// cache don't contains, queue up to size?
		if p.UseQu.Size() > p.MaxSize() {
			if pn := p.UseQu.Pop(); pn != nil {
				delete(p.Addr, pn)
			}
		}
	}
}









