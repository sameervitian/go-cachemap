[![GoDoc](https://godoc.org/github.com/sameervitian/go-cachemap?status.svg)](https://godoc.org/github.com/sameervitian/go-cachemap)
# go-cachemap
Golang local map cache

go-cachemap is a memory cache store implemeNted using inbuilt map. Cache can be used by multiple go routine as it is thread safe. go-cachemap is not intended to be a persistent cache store.


## Installation

go get github.com/sameervitian/go-cachemap

## Usage
```go
package main

import (
	"log"
	"time"

	cache "github.com/sameervitian/go-cachemap"
)

func main() {
	option := cache.Option{
		TTL: 10, // in seconds
	}
	cacheStore := cache.New(&option)

	// Case 1
	cacheObj := cacheStore.NewCacheObject("foo")
	cacheObj.Set("bar")

	val, found := cacheObj.Get()
	//`cache hit. value is bar` will be logged
	if found {
		log.Printf("cache hit. value is %v", val.(string))
	} else {
		log.Printf("cache miss")
	}

	time.Sleep(time.Second * 11)
	val, found = cacheObj.Get()

	// `cache miss` will be logged
	if found {
		log.Printf("cache hit. value is %v", val.(string))
	} else {
		log.Printf("cache miss")
	}

	type Item struct {
		Name  string
		Price int64
	}

	// Case 2
	item := Item{Name: "foo", Price: 100}
	cacheObj = cacheStore.NewCacheObject("foo")
	cacheObj.Set(&item)
	val, found = cacheObj.Get()
	//` cache hit. value is main.Item{Name:"foo", Price:100}` will be logged
	if found {
		log.Printf("cache hit. value is %#v", *(val.(*Item)))
	} else {
		log.Printf("cache miss")
	}

	time.Sleep(time.Second * 2) //Sleep for 2 seconds
	val, found = cacheObj.Get()
	// ` cache hit. value is main.Item{Name:"foo", Price:100}` will be logged
	if found {
		log.Printf("cache hit. value is %#v", *(val.(*Item)))
	} else {
		log.Printf("cache miss")
	}

	time.Sleep(time.Second * 9) //Sleep for 9 seconds

	val, found = cacheObj.Get()
	// `cache miss` will be logged
	if found {
		log.Printf("cache hit. value is %#v", val.(*Item))
	} else {
		log.Printf("cache miss")
	}

	// Case 3
	item = Item{Name: "foo", Price: 100}
	cacheObj = cacheStore.NewCacheObject("foo", cache.CacheObjectOption{TTL: 2}) // overriding ttl to 2 seconds
	cacheObj.Set(&item)
	val, found = cacheObj.Get()
	//`cache hit. value is main.Item{Name:"foo", Price:100}` will be logged
	if found {
		log.Printf("cache hit. value is %#v", *(val.(*Item)))
	} else {
		log.Printf("cache miss")
	}
	time.Sleep(time.Second * 3) //Sleep for 3 seconds

	val, found = cacheObj.Get()
	// `cache miss` will be logged
	if found {
		log.Printf("cache hit. value is %#v", *(val.(*Item)))
	} else {
		log.Printf("cache miss")
	}

}
```
