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

func traverse(self *Node, result *[]int) {
	if self == nil {
		return
	}

	*result = append(*result, self.Val)
	for _, child := range self.Children {
		traverse(child, result)
	}
}
