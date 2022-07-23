package main

import "context"

// Cache interface. If a custom cache is provided, it must implement this interface
type Cache[K comparable, V any] interface {
	Get(context.Context, K) (any, bool)
	Set(context.Context, K, V)
	Delete(context.Context, K) bool
	Clear()
}

// NoCache implement Cache interface where all methods are noops
type NoCache[K comparable, V any] struct{}

// NewNoCache ...
func NewNoCache[K comparable, V any]() Cache[K, V] {
	return &NoCache[K, V]{}
}

// Clear is a NOOP
func (*NoCache[K, V]) Clear() {
}

// Delete is a NOOP
func (*NoCache[K, V]) Delete(context.Context, K) bool {
	return false
}

// Get is a NOOP
func (*NoCache[K, V]) Get(context.Context, K) (any, bool) {
	return nil, false
}

// Set is a NOOP
func (*NoCache[K, V]) Set(context.Context, K, V) {
}
