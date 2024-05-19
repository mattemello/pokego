package main

import "time"

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	cache map[string]cacheEntry
}

func (cach Cache) NewCach(interval time.Duration) {

	cach.realLoop()

	ticker := time.NewTicker(interval.Seconds())

	defer ticker.Stop()
	//go func

}

func (cach Cache) realLoop() {
	for k := range cach.cache {
		delete(cach.cache, k)
	}
}

func (cach Cache) Add(key string, value []byte) {

	cach.cache[key] = cacheEntry{time.Now(), value}

}
