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
