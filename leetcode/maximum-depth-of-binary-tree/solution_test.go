package maximum_depth_of_binary_tree

import "testing"

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		expect int
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
			expect: 3,
		},
	}

	for _, test := range tests {
		if maxDepth(test.root) != test.expect {
			t.Fatal("failed")
		}
	}

}
