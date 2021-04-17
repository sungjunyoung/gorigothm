package delete_node_in_a_linked_list

import "testing"

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		node   *ListNode
		expect *ListNode
	}{
		{
			node: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
			expect: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}

	for _, test := range tests {
		deleteNode(test.node.Next)

		curRes := test.node
		curExp := test.expect
		for curRes != nil {
			if curRes.Val != curExp.Val {
				t.Fatal("failed")
			}
			curRes = curRes.Next
			curExp = curExp.Next
		}
	}
}
