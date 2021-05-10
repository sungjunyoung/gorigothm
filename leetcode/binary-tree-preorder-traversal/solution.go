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

	*res = append(*res, cur.Val)
	traverse(cur.Left, res)
	traverse(cur.Right, res)
}
