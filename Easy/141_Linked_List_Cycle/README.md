# [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle)

## Solution 1: Hash Set
- **Time Complexity**: O(n) - where n is the number of nodes in the linked list, as we need to visit each node at most once
- **Space Complexity**: O(n) - in the worst case, we store all nodes in the hash set if there's no cycle
- **Approach**:
  1. Use a hash set to keep track of visited nodes
  2. Traverse the linked list, for each node:
     - Check if the current node has been visited before
     - If yes, a cycle exists, return true
     - If not, mark the node as visited and move to the next node
  3. If we reach the end of the list (null), return false as no cycle exists


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
- **Space Complexity**: O(1) - only using two pointer variables (slow and fast)
- **Approach**:
  1. Initialize two pointers, slow and fast, both pointing to the head
  2. Move slow pointer one step at a time and fast pointer two steps at a time
  3. If there is a cycle, the fast pointer will eventually catch up to the slow pointer
  4. If there is no cycle, the fast pointer will reach the end of the list
  5. Return true if the pointers meet, false otherwise


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