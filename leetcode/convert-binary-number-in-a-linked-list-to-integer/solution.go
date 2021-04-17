package convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	current := head
	var arr []int
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}

	result := 0
	mul := 1
	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i] * mul
		mul *= 2
	}

	return result
}
