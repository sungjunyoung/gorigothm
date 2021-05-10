package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	result := [][]int{{root.Val}}

	for len(q) > 0 {
		qLen := len(q)

		subResult := []int{}
		for i := 0; i < qLen; i++ {
			cur := q[0]
			q = q[1:]

			if cur.Left != nil {
				q = append(q, cur.Left)
				subResult = append(subResult, cur.Left.Val)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
				subResult = append(subResult, cur.Right.Val)
			}
		}
		if len(subResult) != 0 {
			result = append(result, subResult)
		}
	}

	return result
}
