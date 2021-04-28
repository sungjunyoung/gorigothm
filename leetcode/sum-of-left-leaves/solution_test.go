package sum_of_left_leaves

import "testing"

func Test_sumOfLeftLeaves(t *testing.T) {
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
			expect: 24,
		},
	}

	for _, test := range tests {
		if sumOfLeftLeaves(test.root) != test.expect {
			t.Fatal("failed")
		}
	}
}
