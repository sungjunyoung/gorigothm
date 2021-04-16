package lru_cache

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	Capacity int
	Head     *Node
	Tail     *Node
	Map      map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Map:      map[int]*Node{},
	}
}

func (lc *LRUCache) Get(key int) int {
	node, ok := lc.Map[key]
	if !ok {
		return -1
	}

	lc.remove(node)
	lc.add(node)
	return node.Val
}

func (lc *LRUCache) Put(key int, val int) {
	node, ok := lc.Map[key]
	if ok {
		node.Val = val
		lc.remove(node)
		lc.add(node)
		return
	}

	node = &Node{
		Key: key,
		Val: val,
	}

	if lc.Capacity <= len(lc.Map) {
		delete(lc.Map, lc.Head.Key)
		lc.remove(lc.Head)
	}

	lc.add(node)
	lc.Map[key] = node
	return
}

func (lc *LRUCache) remove(node *Node) {
	if lc.Head == node && lc.Tail == node {
		lc.Head = nil
		lc.Tail = nil
		return
	}

	if lc.Head == node {
		lc.Head = node.Next
	} else if lc.Tail == node {
		lc.Tail = node.Prev
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
}

func (lc *LRUCache) add(node *Node) {
	if lc.Tail == nil {
		lc.Head = node
		lc.Tail = node
		return
	}

	node.Prev = lc.Tail
	lc.Tail.Next = node
	lc.Tail = node
}
