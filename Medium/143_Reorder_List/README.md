# [Reorder List](https://leetcode.com/problems/reorder-list/)

## Solution 1: Tortoise and Hare Algorithm with List Merging
- **Time Complexity**: O(n) - where n is the number of nodes in the linked list
  - Finding the middle takes O(n/2) time
  - Reversing the second half takes O(n/2) time
  - Merging the two halves takes O(n/2) time
  - Overall, O(n) time complexity
- **Space Complexity**: O(1) - constant extra space
  - We only use a few pointer variables regardless of input size
  - All operations are performed in-place without additional data structures
- **Approach**:
  1. Find the middle of the linked list using slow and fast pointers
  2. Split the list into two halves
  3. Reverse the second half
  4. Merge the two halves by alternating nodes
- **Key Insights**:
  - Fast and slow pointers technique is crucial for finding the middle point of a linked list in a single pass
  - In-place reversal of a linked list => no extra place


```go
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	second := reverse(slow.Next)
	slow.Next = nil // Cut first half

	// Merge two halves
	first := head
	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next

		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

```


