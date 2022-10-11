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

type LinkedMap[K comparable, V any] struct {
	m        map[K]elem[K, V]
	lastKey  K
	startKey K
}

type elem[K comparable, V any] struct {
	v    V
	next K
}

func (l *LinkedMap[K, V]) Get(k K) (V, bool) {
	if vv, ok := l.m[k]; ok {
		return vv.v, ok
	}
	var v V
	return v, false
}

func (l *LinkedMap[K, V]) Put(k K, v V) {
	if len(l.m) == 0 {
		l.m = make(map[K]elem[K, V])
		l.startKey = k
	}
	l.m[k] = elem[K, V]{
		v: v,
	}
	if _, ok := l.m[l.lastKey]; ok {
		l.m[l.lastKey] = elem[K, V]{
			v:    l.m[l.lastKey].v,
			next: k,
		}
	}
	l.lastKey = k
}

func (l *LinkedMap[K, V]) Range(f func(k K, v V)) {
	if v, ok := l.m[l.startKey]; ok {
		f(l.startKey, v.v)
		l.gGet(v, f)
	}
}

func (l *LinkedMap[K, V]) gGet(val elem[K, V], f func(k K, v V)) {
	if _, ok := l.m[val.next]; ok {
		f(val.next, l.m[val.next].v)
		l.gGet(l.m[val.next], f)
	}
}
