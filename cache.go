package cachestore

import (
	"log"
	"sync"
	"time"
)

type Store struct {
	ttl   int64
	cache map[string]interface{}
	sync.RWMutex
}

type Option struct {
	TTL int64
}

func New(o *Option) *Store {
	return &Store{ttl: o.TTL, cache: make(map[string]interface{}, 0)}
}

type CacheObject struct {
	key   string
	store *Store
	ttl   int64
}

type CacheObjectOption struct {
	TTL int64
}

func (s *Store) NewCacheObject(key string, option ...CacheObjectOption) *CacheObject {
	ttl := s.ttl
	if len(option) > 0 {
		ttl = option[0].TTL
	}
	return &CacheObject{key: key, store: s, ttl: ttl}
}

func (c *CacheObject) Get() (interface{}, bool) {
	c.store.Lock()
	defer c.store.Unlock()

	value, ok := c.store.cache[c.key]
	return value, ok
}

func (c *CacheObject) Set(val interface{}) {
	c.store.Lock()
	defer c.store.Unlock()
	c.store.cache[c.key] = val
	c.setttl()
}

func (c *CacheObject) Expire() {
	c.store.Lock()
	defer c.store.Unlock()

	delete(c.store.cache, c.key)

}

func (c *CacheObject) setttl() {
	t := time.NewTicker(time.Second * time.Duration(c.ttl))
	go func() {
		<-t.C
		c.Expire()
		log.Printf("key `%v` evicted after %vs", c.key, c.ttl)
	}()
}
