package main_test

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arr := [101]int{}
	for i := range arr {
		arr[i] = -1
	}

	count := 0
	node1 := l1
	node2 := l2
	for {
		if node1 == nil && node2 == nil {
			break
		}

		tmp1 := 0
		tmp2 := 0
		if node1 != nil {
			tmp1 = node1.Val
		}
		if node2 != nil {
			tmp2 = node2.Val
		}

		arr[count] = tmp1 + tmp2
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}

		count += 1
	}

	count = 0
	for {
		if arr[count] == -1 {
			break
		}

		if arr[count] >= 10 {
			arr[count] %= 10
			if arr[count+1] == -1 {
				arr[count+1] = 1
			} else {
				arr[count+1] += 1
			}
		}
		count += 1
	}

	result := &ListNode{
		Val:  arr[0],
		Next: nil,
	}
	next := result
	count = 1
	for {
		if arr[count] == -1 {
			break
		}

		next.Next = &ListNode{
			Val:  arr[count],
			Next: nil,
		}
		next = next.Next
		count += 1
	}

	return result
}

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
