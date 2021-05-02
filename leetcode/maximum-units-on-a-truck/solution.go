package maximum_units_on_a_truck

import (
	"container/heap"
)

type boxList []int

func (bl boxList) Len() int {
	return len(bl)
}

func (bl boxList) Less(i, j int) bool {
	return bl[i] > bl[j]
}

func (bl boxList) Swap(i, j int) {
	bl[i], bl[j] = bl[j], bl[i]
}

func (bl *boxList) Push(x interface{}) {
	*bl = append(*bl, x.(int))
}

func (bl *boxList) Pop() interface{} {
	old := *bl
	n := len(old)
	elem := old[n-1]
	*bl = old[:n-1]
	return elem
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	boxs := &boxList{}
	for i := range boxTypes {
		nums := boxTypes[i][0]
		for j := 0; j < nums; j++ {
			heap.Push(boxs, boxTypes[i][1])
		}
	}

	result := 0
	for i := 0; i < truckSize; i++ {
		if boxs.Len() == 0 {
			return result
		}
		result += heap.Pop(boxs).(int)
	}

	return result
}
