package main

import (
	"github.com/allegro/bigcache"
	"log"
	"time"
)

func main() {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Println(1111111111)
		log.Println(err)
		return
	}

	entry, err := cache.Get("my-unique-key")
	if err != nil {
		log.Println(22222222)
		log.Println(err)
		log.Println(22222223)
		return
	}

	if entry == nil {
		entry = []byte("value")
		cache.Set("my-unique-key", entry)
		log.Println(333333)
	}

	log.Println(string(entry))
}
