# [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list)

## Solution 1: Iterative Approach
- **Time Complexity**: O(n) - where n is the number of nodes in the linked list
- **Space Complexity**: O(1) - only using a constant amount of extra space (three variables: pre, cur, and temp)
- **Approach**:
  1. Initialize two pointers: `pre` pointing to `nil` and `cur` pointing to the head
  2. While `cur` is not `nil`:
     - Store the next node temporarily in `temp`
     - Reverse the current node's pointer to point to the previous node
     - Move the previous pointer to the current node
     - Move the current pointer to the temporary next node
  3. Return the new head (the `pre` pointer)


```go
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

```