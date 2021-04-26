package binary_tree_inorder_traversal

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		expect []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expect: []int{1, 3, 2},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(inorderTraversal(test.root), test.expect) {
			t.Fatal("failed")
		}
	}
}
