package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricBfs(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

	q := []*TreeNode{}
	q = append(q, root)

	for len(q) != 0 {
		qLen := len(q)

		checks := []int{}
		for i := 0; i < qLen; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
				checks = append(checks, cur.Left.Val)
			} else {
				checks = append(checks, -1)
			}

			if cur.Right != nil {
				q = append(q, cur.Right)
				checks = append(checks, cur.Right.Val)
			} else {
				checks = append(checks, -1)
			}

		}

		for i := 0; i < len(checks)/2; i++ {
			if checks[i] != checks[len(checks)-1-i] {
				return false
			}
		}

	}
	return true
}

func isSymmetricDfs(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
