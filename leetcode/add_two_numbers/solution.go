package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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
