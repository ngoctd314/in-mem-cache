package main

import (
	"context"
	"sync"
)

// PerRequestCache is an in mem implementation of Cache interface
// This implementation is well suited for a "per-request" cache
// It's not well suited for long lived cached items
type PerRequestCache[K comparable, V any] struct {
	items map[K]V
	mu    sync.RWMutex
}

// NewPerRequestCache ...
func NewPerRequestCache[K comparable, V any]() Cache[K, V] {
	items := make(map[K]V)

	return &PerRequestCache[K, V]{
		items: items,
	}
}

// Clear implements Cache
func (cache *PerRequestCache[K, V]) Clear() {
	cache.mu.Lock()
	cache.items = make(map[K]V)
	cache.mu.Unlock()
}

// Delete implements Cache
func (cache *PerRequestCache[K, V]) Delete(ctx context.Context, key K) bool {
	if _, found := cache.Get(ctx, key); found {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		delete(cache.items, key)
		return true
	}
	return false
}

// Get implements Cache
func (cache *PerRequestCache[K, V]) Get(_ context.Context, key K) (any, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	if item, found := cache.items[key]; found {
		return item, true
	}

	return nil, false
}

// Set implements Cache
func (cache *PerRequestCache[K, V]) Set(_ context.Context, key K, value V) {
	cache.mu.Lock()
	cache.items[key] = value
	cache.mu.Unlock()
}
