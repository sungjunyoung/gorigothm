package dominator

type info struct {
	index int
	count int
}

func Solution(A []int) int {
	infoMap := map[int]*info{}

	for i, a := range A {
		if _, ok := infoMap[a]; ok {
			infoMap[a].count += 1
		} else {
			infoMap[a] = &info{
				index: i,
				count: 1,
			}
		}
	}

	for _, v := range infoMap {
		if v.count > len(A)/2 {
			return v.index
		}
	}

	return -1
}
