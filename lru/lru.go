package lru

import (
	"errors"
	"fmt"
	"github.com/Golang_play/list"
)

type evictcb func(key interface{}, val interface{})

type lru struct {
	cachelen  int
	cache     map[interface{}]*list.Element
	evictlist *list.List
	onevict   evictcb
}

type kv struct {
	key interface{}
	val interface{}
}

func (l *lru) Len() int {
	return l.evictlist.Size()
}

func New(cachesize int, cb evictcb) (*lru, error) {
	if cachesize < 0 {
		return nil, errors.New("provide positive")
	}
	l := lru{
		cachelen: cachesize,
		//map of key:value
		cache:     make(map[interface{}]*list.Element),
		evictlist: list.New(),
		onevict:   cb,
	}
	return &l, nil
}

func (l *lru) Add(key interface{}, val interface{}) bool {
	if val1, ok := l.cache[key]; ok {
		l.evictlist.MoveToFront(val1)
		return false
	} else {
		n := &kv{key, val}
		l.cache[key] = l.evictlist.PushToFront(n)
	}

	if l.evictlist.Size() > l.cachelen {
		l.RemoveOldest()
		return true
	}
	return false
}

func (l *lru) RemoveOldest() {
	n := l.evictlist.Back()
	if n != nil {
		l.evictlist.Remove(n)

		kv := n.Val.(*kv)
		delete(l.cache, kv.key)
		if l.onevict != nil {
			l.onevict(kv.key, kv.val)
		}
	}
}

func (l *lru) remove(key interface{}) {

}

//func (l *lru) Get(key interface{}) {

func (l *lru) purge() {
	for k, v := range l.cache {
		delete(l.cache, k)
		fmt.Println("EVICT:", l.evictlist.Remove(v).(*kv).val.(int))
	}
}
