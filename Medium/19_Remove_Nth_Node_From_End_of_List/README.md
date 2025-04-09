# [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## Solution 1: Sorting Approach
- **Time Complexity**: O(n) - where n is the length of the linked list, as we need to traverse the list twice
- **Space Complexity**: O(1) - using only constant extra space for pointers regardless of input size
- **Approach**:
  1. Create a dummy node that points to the head to handle edge cases (like removing the head node)
  2. First pass: Count the total number of nodes in the linked list
  3. Calculate the position of the node to be removed from the beginning
  4. Second pass: Traverse to the node just before the one to be removed
  5. Skip the target node by updating the next pointer to point to the node after the target
  6. Return the new head (dummy.Next)


```go
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
```