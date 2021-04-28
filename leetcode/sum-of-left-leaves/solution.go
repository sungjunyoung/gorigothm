package sum_of_left_leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	//   3
	//  / \
	// 9  20
	//   /  \
	//  15   7

	var q []*TreeNode
	q = append(q, root)

	result := 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		if cur.Left != nil {
			q = append(q, cur.Left)

			if cur.Left.Left == nil && cur.Left.Right == nil {
				result += cur.Left.Val
			}
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}

	return result
}
