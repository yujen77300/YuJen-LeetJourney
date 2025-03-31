# [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list)

## Solution 1: Hash Set
- **Time Complexity**: O(n) - where n is the number of nodes in the linked list
- **Space Complexity**: O(1) - only using two pointer variables (slow and fast)
- **Approach**:
  1. Initialize two pointers, slow and fast, both pointing to the head
  2. Move slow pointer one step at a time and fast pointer two steps at a time
  3. If there is a cycle, the fast pointer will eventually catch up to the slow pointer
  4. If there is no cycle, the fast pointer will reach the end of the list
  5. Return true if the pointers meet, false otherwise


```go
func hasCycle(head *ListNode) bool {
    visited := make(map[*ListNode]bool)

    for head != nil {
        if visited[head] {
            return true
        }
        visited[head] = true
        head = head.Next
    }

    return false
}

```


## Solution 2: Floyd Cycle Detection Algorithm
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
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
```