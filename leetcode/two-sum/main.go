package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%#v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%#v\n", twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, current := range nums {
		match := target - current
		if _, ok := hash[match]; ok {
			return []int{i, hash[match]}
		}

		hash[current] = i
	}
	return []int{}
}
