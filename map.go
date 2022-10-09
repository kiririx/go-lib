package go_lib

import "sync"

type LockMap[K comparable, V any] struct {
	lock sync.RWMutex
	m    map[K]V
}

func (m *LockMap[K, V]) Put(k K, v V) {
	m.lock.Lock()
	m.m[k] = v
	m.lock.Unlock()
}

func (m *LockMap[K, V]) Get(k K) V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.m[k]
}

func (m *LockMap[K, V]) Range(f func(k K, v V)) {
	m.lock.RLock()
	for key, val := range m.m {
		f(key, val)
	}
	m.lock.RUnlock()
}
