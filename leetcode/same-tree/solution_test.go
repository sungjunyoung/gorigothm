package same_tree

import "testing"

func Test_isSameTree(t *testing.T) {
	tests := []struct {
		p      *TreeNode
		q      *TreeNode
		expect bool
	}{
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expect: true,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			expect: false,
		},
	}

	for _, test := range tests {
		if isSameTree(test.p, test.q) != test.expect {
			t.Error("failed")
		}
	}
}
