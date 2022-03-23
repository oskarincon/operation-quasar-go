package cache

import (
	"fmt"
	"time"

	"github.com/oskarincon/operation-quasar-go/models"
	"github.com/patrickmn/go-cache"
)

var (
	DefaultMemCache Cache
)

type MemCache struct {
	cli *cache.Cache
}

func init() {
	fmt.Printf("[cache] - init")
	cli := cache.New(10*time.Minute, 10*time.Minute)
	DefaultMemCache = &MemCache{
		cli: cli,
	}
}

func (mem *MemCache) Get(key string) (val interface{}, exist bool) {
	fmt.Printf("[cache] - Get key: %#v\n", key)
	valI, exist := mem.cli.Get(key)
	fmt.Printf("[cache] - valI key: %#v\n, exist: %#v\n", valI, exist)
	if exist == true {
		val = valI.(models.Satellite)
	}
	return
}

func (mem *MemCache) Set(key string, val interface{}, expiration time.Duration) {
	fmt.Printf("[cache] - Set key: %#v\n, data: %#v\n", key, val)
	mem.cli.Set(key, val, expiration)
}
