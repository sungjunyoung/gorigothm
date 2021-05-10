package binary_tree_level_order_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder(t *testing.T) {
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
			expect: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, levelOrder(test.root), test.expect)
	}
}
