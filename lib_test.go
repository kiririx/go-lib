package go_lib

import "testing"

func TestSet(t *testing.T) {
	s := Set[string]{}
	s.Put("v")
	s.Put("k")
	s.Range(func(k string) {
		t.Log(k)
	})
}

func TestLinkMap(t *testing.T) {
	lm := LinkedMap[string, string]{}
	lm.Put("a", "1")
	lm.Put("b", "1")
	lm.Put("z", "1")
	lm.Put("g", "1")

	t.Log(lm.Get("a"))

	lm.Range(func(k string, v string) {
		t.Log(k, v)
	})
}

func TestMap(t *testing.T) {
	m := map[string]string{}
	m["a"] = "1"
	m["b"] = "1"
	m["z"] = "1"
	m["g"] = "1"

	for k, v := range m {
		t.Log(k, v)
	}
}

func TestSyncMap(t *testing.T) {
	sm := SyncMap[string, int]{}
	sm.Store("abc", 2)
	t.Log(sm.Load("abc"))
	sm.Range(func(k string, v int) bool {
		t.Log(k, v)
		return true
	})
}
