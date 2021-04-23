package merge_two_sorted_lists

import "testing"

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		l1     *ListNode
		l2     *ListNode
		expect *ListNode
	}{
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			expect: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := mergeTwoLists(test.l1, test.l2)
		for result != nil {
			if test.expect.Val != result.Val {
				t.Fatal("failed")
			}

			test.expect = test.expect.Next
			result = result.Next
		}
	}
}
