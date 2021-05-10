package binary_tree_preorder_traversal

import (
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, preorderTraversal(test.root), test.expect)
	}
}
