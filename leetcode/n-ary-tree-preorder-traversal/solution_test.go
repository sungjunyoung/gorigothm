package n_ary_tree_preorder_traversal

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	tests := []struct {
		root   *Node
		expect []int
	}{
		{
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 3,
						Children: []*Node{
							{
								Val: 5,
							},
							{
								Val: 6,
							},
						},
					},
					{
						Val: 2,
					},
					{
						Val: 4,
					},
				},
			},
			expect: []int{1, 3, 5, 6, 2, 4},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(preorder(test.root), test.expect) {
			t.Fatal("failed")
		}
	}
}
