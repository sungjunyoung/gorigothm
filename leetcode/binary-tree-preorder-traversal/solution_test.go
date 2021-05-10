package binary_tree_preorder_traversal

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
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
			expect: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(preorderTraversal(test.root), test.expect) {
			t.Error("failed")
		}
	}
}
