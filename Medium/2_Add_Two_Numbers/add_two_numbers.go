package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		total := val1 + val2 + carry
		carry = total / 10
		val := total % 10

		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return dummy.Next
}
