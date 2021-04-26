package binary_tree_level_order_traversal_ii

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		expect [][]int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expect: [][]int{{15, 7}, {9, 20}, {3}},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(levelOrderBottom(test.root), test.expect) {
			t.Fatal("failed")
		}
	}
}
