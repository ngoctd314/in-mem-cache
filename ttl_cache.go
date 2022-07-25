package main

import "time"

// Item ..
type Item struct {
	Object     any
	Expiration int64
}

// Expired ...
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}

	return time.Now().UnixNano() > item.Expiration
}

// Expiration type
const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

// TTLCache ...
type TTLCache struct{}

// type cache[K comparable, V any] struct {
// 	defaultExpiration time.Duration
// 	items             map[K]V
// }
