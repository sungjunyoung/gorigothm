package add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		l1     *ListNode
		l2     *ListNode
		expect *ListNode
	}{
		{
			l1: &ListNode{
				Val: 2, Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 3, Next: nil,
					},
				},
			},
			l2: &ListNode{
				Val: 5, Next: &ListNode{
					Val: 6, Next: &ListNode{
						Val: 4, Next: nil,
					},
				},
			},
			expect: &ListNode{
				Val: 7, Next: &ListNode{
					Val: 0, Next: &ListNode{
						Val: 8, Next: nil,
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val:  0,
				Next: nil,
			},
			l2: &ListNode{
				Val:  0,
				Next: nil,
			},
			expect: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			l1: &ListNode{
				Val: 9, Next: &ListNode{
					Val: 9, Next: &ListNode{
						Val: 9, Next: &ListNode{
							Val: 9, Next: &ListNode{
								Val: 9, Next: &ListNode{
									Val: 9, Next: &ListNode{
										Val: 9, Next: nil,
									},
								},
							},
						},
					},
				},
			},
			l2: &ListNode{
				Val: 9, Next: &ListNode{
					Val: 9, Next: &ListNode{
						Val: 9, Next: &ListNode{
							Val: 9, Next: nil,
						},
					},
				},
			},
			expect: &ListNode{
				Val: 8, Next: &ListNode{
					Val: 9, Next: &ListNode{
						Val: 9, Next: &ListNode{
							Val: 9, Next: &ListNode{
								Val: 0, Next: &ListNode{
									Val: 0, Next: &ListNode{
										Val: 0, Next: &ListNode{
											Val: 1, Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := addTwoNumbers(test.l1, test.l2)
		if !reflect.DeepEqual(result, test.expect) {
			t.Fatal("failed")
		}
	}
}
