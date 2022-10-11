package go_lib

import "fmt"

type Set[K comparable] struct {
	m map[K]bool
}

func (s *Set[K]) Put(k K) {
	if s.m == nil {
		s.m = make(map[K]bool)
	}
	s.m[k] = true
}

func (s *Set[K]) Exist(k K) bool {
	if _, ok := s.m[k]; ok {
		return true
	}
	return false
}

func (s *Set[K]) Range(f func(K)) {
	for k, _ := range s.m {
		f(k)
	}
}

func (s *Set[K]) String() string {
	str := "["
	s.Range(func(k K) {
		str += fmt.Sprintf("%v ", k)
	})
	str += "]"
	return str
}

type LockSet[K comparable] struct {
	m LockMap[K, bool]
}

func (s *LockSet[K]) Put(k K) {
	if s.m.m == nil {
		s.m.m = make(map[K]bool)
	}
	s.m.m[k] = true
}

func (s *LockSet[K]) Exist(k K) bool {
	if _, ok := s.m.m[k]; ok {
		return true
	}
	return false
}

func (s *LockSet[K]) Range(f func(K)) {
	for k, _ := range s.m.m {
		f(k)
	}
}

func (s *LockSet[K]) String() string {
	str := "["
	s.Range(func(k K) {
		str += fmt.Sprintf("%v ", k)
	})
	str += "]"
	return str
}

type LinkSet[K comparable] struct {
	m LinkedMap[K, bool]
}

func (l *LinkSet[K]) Exist(k K) bool {
	if _, ok := l.m.Get(k); ok {
		return true
	}
	return false
}

func (l *LinkSet[K]) Put(k K) {
	l.m.Put(k, true)
}

func (l *LinkSet[K]) Range(f func(k K)) {
	l.m.Range(func(k K, _ bool) {
		f(k)
	})
}
