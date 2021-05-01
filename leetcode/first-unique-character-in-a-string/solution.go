package first_unique_character_in_a_string

import "sort"

type counter struct {
	count int
	index int
}

func firstUniqChar(s string) int {
	countMap := map[int32]*counter{}

	for i, c := range s {
		if countMap[c] == nil {
			countMap[c] = &counter{
				count: 1,
				index: i,
			}
		} else {
			countMap[c].count++
		}
	}

	results := []int{}
	for _, c := range countMap {
		if c.count == 1 {
			results = append(results, c.index)
		}
	}

	if len(results) == 0 {
		return -1
	}

	sort.Ints(results)
	return results[0]
}
