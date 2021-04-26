package binary_tree_level_order_traversal_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	var results [][]int
	if root == nil {
		return results
	}

	for len(queue) != 0 {
		var result []int

		size := len(queue)
		for i := 0; i < size; i++ {
			now := queue[0]
			queue = queue[1:]
			result = append(result, now.Val)
			if now.Left != nil {
				queue = append(queue, now.Left)
			}
			if now.Right != nil {
				queue = append(queue, now.Right)
			}
		}

		results = append([][]int{result}, results[0:]...)
	}

	return results
}
