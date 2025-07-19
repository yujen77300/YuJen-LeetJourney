# [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

## Solution 1: Simulate Addition with Carry
- **Time Complexity**: O(max(m, n)) - where m and n are the lengths of the two linked lists
  - Traverse each linked list exactly once
  - The loop continues until both lists are exhausted and no carry remains
- **Space Complexity**: O(max(m, n)) - for the result linked list
  - Create a new linked list to store the result
  - The result length is at most max(m, n) + 1 (due to potential final carry)
- **Approach**:
  1. Use a dummy head node to simplify linked list construction
  2. Traverse both linked lists simultaneously from head to tail
  3. At each position, add the corresponding digits plus any carry from previous addition
  4. Handle carry by dividing the sum by 10, store remainder as current digit
  5. Continue until both lists are exhausted and no carry remains
  6. Return the next node of dummy head (actual result)
- **Key Insights**:
  - Dummy head technique simplifies linked list construction and edge case handling
  - Don't forget to handle the final carry after both lists are exhausted


```go
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

```


