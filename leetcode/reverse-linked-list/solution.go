package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	var prevPrevNode *ListNode = nil
	prevNode := head
	currentNode := head.Next
	for currentNode != nil {
		prevNode.Next = prevPrevNode
		prevPrevNode = prevNode
		prevNode = currentNode
		currentNode = currentNode.Next
	}

	prevNode.Next = prevPrevNode
	return prevNode
}
