package range_sum_of_bst

import "testing"

func Test_rangeSumBST(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		low    int
		high   int
		expect int
	}{
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 15,
					Right: &TreeNode{
						Val: 18,
					},
				},
			},
			low:    7,
			high:   15,
			expect: 32,
		},
	}

	for _, test := range tests {
		if rangeSumBST(test.root, test.low, test.high) != test.expect {
			t.Fatalf("result %d is not expected %d", rangeSumBST(test.root, test.low, test.high), test.expect)
		}
	}
}
