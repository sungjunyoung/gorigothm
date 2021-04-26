package minimum_absolute_difference_in_bst

import "testing"

func Test_getMinimumDifference(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		expect int
	}{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			expect: 1,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 48,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 49,
					},
				},
			},
			expect: 1,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			expect: 1,
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			expect: 1,
		},
	}

	for _, test := range tests {
		if getMinimumDifference(test.root) != test.expect {
			t.Fatal("failed")
		}
	}
}
