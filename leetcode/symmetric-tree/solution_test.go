package symmetric_tree

import (
	"testing"
)

var tests = []struct {
	root   *TreeNode
	expect bool
}{
	{
		root: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 3},
			},
		},
		expect: false,
	},
	{
		root: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 3},
			},
		},
		expect: true,
	},
	{
		root: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 3},
			},
		},
		expect: false,
	},
}

func Test_isSymmetricBfs(t *testing.T) {
	for i, test := range tests {
		if isSymmetricBfs(test.root) != test.expect {
			t.Errorf("failed for test %d", i)
		}
	}
}

func Test_isSymmetricDfs(t *testing.T) {
	for i, test := range tests {
		if isSymmetricDfs(test.root) != test.expect {
			t.Errorf("failed for test %d", i)
		}
	}
}
