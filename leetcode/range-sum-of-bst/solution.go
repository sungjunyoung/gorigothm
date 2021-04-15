package range_sum_of_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	result := 0
	if low <= root.Val && root.Val <= high {
		result = root.Val
	}

	return result + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
