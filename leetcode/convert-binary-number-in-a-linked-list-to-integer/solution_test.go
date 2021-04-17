package convert_binary_number_in_a_linked_list_to_integer

import "testing"

func Test_getDecimalValue(t *testing.T) {
	tests := []struct {
		head   *ListNode
		expect int
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			expect: 5,
		},
	}

	for _, test := range tests {
		if getDecimalValue(test.head) != test.expect {
			t.Fatal("failed")
		}
	}
}
