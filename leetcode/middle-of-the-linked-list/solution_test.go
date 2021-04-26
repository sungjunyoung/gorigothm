package middle_of_the_linked_list

import "testing"

func Test_middleNode(t *testing.T) {
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
								Val: 5,
							},
						},
					},
				},
			},
			expect: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			expect: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	for _, test := range tests {
		result := middleNode(test.head)
		for result != nil && test.expect != nil {
			if result.Val != test.expect.Val {
				t.Fatal("failed")
			}
			result = result.Next
			test.expect = test.expect.Next
		}
	}
}
