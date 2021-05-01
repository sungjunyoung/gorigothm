package invert_binary_tree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		expect *TreeNode
	}{
		{
			root: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expect: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
	}

	for i, test := range tests {
		result := invertTree(test.root)
		resultArr := []int{}
		dfs(result, &resultArr)
		expectArr := []int{}
		dfs(test.expect, &expectArr)

		if !reflect.DeepEqual(resultArr, expectArr) {
			t.Errorf("test %d failed", i)
		}
	}
}

func dfs(root *TreeNode, results *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, results)
	*results = append(*results, root.Val)
	dfs(root.Right, results)
}
