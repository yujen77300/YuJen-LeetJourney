# [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

## Solution 1: Iterative Approach
- **Time Complexity**: O(n + m) - where n and m are the lengths of the two lists
  - Traverse each node of both lists exactly once
  - Each operation inside the loop takes constant time
- **Space Complexity**: O(1) - constant extra space
  - Only a few pointer variables regardless of input size
  - The merged list reuses the nodes from the input lists
- **Approach**:
  1. Create a dummy node to serve as the starting point of the merged list
  2. Use a current pointer to track where to append the next node
  3. Iterate through both lists simultaneously, comparing values
  4. Append the node with the smaller value to the result list
  5. Move the pointer of the list that contributed the node
  6. After exhausting one list, append the remaining part of the other list
  7. Return the next of the dummy node as the head of the merged list


```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next

}

```




## Solution 2: Recursive Approach 
- **Time Complexity**: O(n + m) - where n and m are the lengths of the two lists
  - Make a recursive call for each node in both lists
  - Each recursive call performs constant time operations
- **Space Complexity**: O(n + m) - for the recursion call stack
  - In the worst case, the depth of recursion equals the sum of both list lengths
  - Each recursive call adds a frame to the call stack
- **Approach**:
  1. Handle base cases: if either list is empty, return the other list
  2. Compare the values of the head nodes of both lists
  3. Recursively set the next pointer of the node with the smaller value
  4. Return the head of the merged list (the node with the smaller value)
- **Key Insights**:
  - The recursive approach elegantly handles the merging logic
  - Each recursive call establishes the next connection for the current minimum node


```go

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }

    if list1.Val < list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoLists(list1, list2.Next)
        return list2
    }
}


```


