package design_hashmap

import "testing"

func Test_MyHashMap(t *testing.T) {
	mhm := Constructor()
	mhm.Put(1, 1)
	mhm.Put(2, 2)
	a := mhm.Get(1)
	if a != 1 {
		t.Fatal("failed")
	}
	b := mhm.Get(3)
	if b != -1 {
		t.Fatal("failed")
	}
	mhm.Put(2, 1)
	c := mhm.Get(2)
	if c != 1 {
		t.Fatal("failed")
	}
	mhm.Remove(2)
	d := mhm.Get(2)
	if d != -1 {
		t.Fatal("failed")
	}
}
