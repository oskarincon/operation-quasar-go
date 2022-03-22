package cache

import (
	"time"
)

type Cache interface {
	Get(key string) (val interface{}, exist bool)
	Set(key string, val interface{}, expiration time.Duration)
}
