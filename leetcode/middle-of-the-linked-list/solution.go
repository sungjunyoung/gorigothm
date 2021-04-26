package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	runner := head
	walker := head

	for runner != nil {
		if runner.Next == nil {
			return walker
		}
		if runner.Next.Next == nil {
			return walker.Next
		}

		runner = runner.Next.Next
		walker = walker.Next
	}
	return nil
}
