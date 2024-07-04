/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    dummy := &ListNode{}
	current := dummy
	temp := head

	for temp != nil {
		if temp.Val == 0 {
			sum := 0
			temp = temp.Next
			for temp != nil && temp.Val != 0 {
				sum += temp.Val
				temp = temp.Next
			}

			if sum != 0 {
				current.Next = &ListNode{Val: sum}
				current = current.Next
			}
		} else {
			temp = temp.Next
		}
	}

	return dummy.Next
}