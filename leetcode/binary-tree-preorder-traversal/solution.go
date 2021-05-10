package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	traverse(root, &result)
	return result
}

func traverse(cur *TreeNode, res *[]int) {
	if cur == nil {
		return
	}

	traverse(cur.Left, res)
	*res = append(*res, cur.Val)
	traverse(cur.Right, res)
}
