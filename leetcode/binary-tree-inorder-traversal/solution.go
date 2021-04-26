package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	traverse(root, &result)
	return result
}

func traverse(self *TreeNode, result *[]int) {
	if self == nil {
		return
	}

	traverse(self.Left, result)
	*result = append(*result, self.Val)
	traverse(self.Right, result)
}
