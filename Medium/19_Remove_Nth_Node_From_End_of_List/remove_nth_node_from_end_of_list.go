package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}

    totalNode := 0
    cur := head
    for cur != nil {
        totalNode += 1
        cur = cur.Next
    }

    needToRemove := totalNode - n
    cur = dummy

    for i := 0; i < needToRemove; i++ {
        cur = cur.Next
    }

    cur.Next = cur.Next.Next

    return dummy.Next
}
