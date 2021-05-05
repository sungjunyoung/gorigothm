package lru_cache

import (
	"testing"
)

func TestLRUCache_GetPut1(t *testing.T) {
	lc := Constructor(2)
	lc.Put(1, 1) // cache is {1=1}
	lc.Put(2, 2) // cache is {1=1, 2=2}
	lc.Get(1)    // return 1
	lc.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lc.Get(2)    // returns -1 (not found)
	lc.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	lc.Get(1)    // return -1 (not found)
	lc.Get(3)    // return 3
	lc.Get(4)    // return 4
}

func TestLRUCache_GetPut2(t *testing.T) {
	lc := Constructor(1)
	lc.Put(2, 1) // cache is {1=1}
	lc.Get(2)    // return 1
	lc.Put(3, 2) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lc.Get(2)    // returns -1 (not found)
	lc.Get(3)    // returns -1 (not found)
}
