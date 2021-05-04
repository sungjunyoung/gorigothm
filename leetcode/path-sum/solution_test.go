package path_sum

import "testing"

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		root      *TreeNode
		targetSum int
		expect    bool
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			targetSum: 4,
			expect:    true,
		},
	}

	for i, test := range tests {
		if hasPathSum(test.root, test.targetSum) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
