package reverse_linked_list

import "testing"

func Test_reverseList(t *testing.T) {
	tests := []struct {
		head   *ListNode
		expect *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			expect: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			expect: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	for _, test := range tests {
		if reverseList(test.head).Val != test.expect.Val {
			t.Fatal("failed")
		}
	}
}
