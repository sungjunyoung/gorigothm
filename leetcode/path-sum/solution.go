package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return traverse(root, targetSum, 0)
}

func traverse(root *TreeNode, targetSum int, curVal int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return curVal+root.Val == targetSum
	}

	left := traverse(root.Left, targetSum, curVal+root.Val)
	right := traverse(root.Right, targetSum, curVal+root.Val)
	return left || right
}
