package n_ary_tree_preorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	traverse(root, &result)
	return result
}

func traverse(root *Node, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	if root.Children != nil {
		for _, child := range root.Children {
			traverse(child, result)
		}
	}
}
