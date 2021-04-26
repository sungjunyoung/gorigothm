package minimum_absolute_difference_in_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	prev := -1
	result := math.MaxInt64
	traverse(&prev, root, &result)
	return result
}

func traverse(prev *int, self *TreeNode, min *int) {
	if self == nil {
		return
	}

	traverse(prev, self.Left, min)

	if *prev == -1 {
		*prev = self.Val
	} else {
		if self.Val-*prev < *min {
			*min = self.Val - *prev
		}
		*prev = self.Val
	}

	traverse(prev, self.Right, min)
}
