package safemap

import (
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{mu: sync.RWMutex{}, m: make(map[K]V)}
}

func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	val, ok := sm.m[key]
	sm.mu.RUnlock()
	return val, ok
}

func (sm *SafeMap[K, V]) Store(key K, val V) {
	sm.mu.Lock()
	sm.m[key] = val
	sm.mu.Unlock()
}

func (sm *SafeMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	delete(sm.m, key)
	sm.mu.Unlock()
}
