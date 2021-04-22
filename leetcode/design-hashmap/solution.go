package design_hashmap

type MyHashMap struct {
	m map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{m: map[int]int{}}
}

func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

func (this *MyHashMap) Get(key int) int {
	val, ok := this.m[key]
	if !ok {
		return -1
	}
	return val
}

func (this *MyHashMap) Remove(key int) {
	delete(this.m, key)
}
