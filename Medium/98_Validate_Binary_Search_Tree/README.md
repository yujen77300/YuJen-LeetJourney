# [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)

## Solution 1: Inorder traversal
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Each node is visited exactly once during the traversal
  - The array comparison is also O(n) in the worst case
- **Space Complexity**: O(n) - for storing:
  - The recursion call stack (up to the height of the tree, which is O(n) in the worst case)
  - The array that holds all node values
- **Approach**:
  1. Perform an inorder traversal of the BST (left-root-right)
  2. Store all visited node values in an array
  3. Check if the resulting array is sorted in strictly increasing order
  4. If any adjacent values are out of order (not strictly increasing), the tree is not a valid BST


> In an inorder traversal a node is visited after its left subtree and before its right subtree


```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	nums := []int{}
	inOrderTraversal(root, &nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}

	return true

}

func inOrderTraversal(node *TreeNode, nums *[]int) {
	if node.Left != nil {
		inOrderTraversal(node.Left, nums)
	}
	*nums = append(*nums, node.Val)
	if node.Right != nil {
		inOrderTraversal(node.Right, nums)
	}
}

```


## Solution 2: Optimized Inorder Traversal with Previous Pointer
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Each node is visited exactly once during the traversal
- **Space Complexity**: O(h) - where h is the height of the tree
  - Only needs recursion call stack space (h), which is O(log n) for balanced trees and O(n) for worst case
  - No additional array is needed to store all values
- **Approach**:
  1. Perform an inorder traversal of the BST (left-root-right)
  2. Instead of storing all values in an array, maintain a "previous node" pointer
  3. During traversal, compare each node's value with the previous node's value
  4. In a valid BST, each node must have a value greater than the previous node
  5. If any node's value is less than or equal to its previous node, the tree is not a valid BST
- **Key Insights**:
  - Using a previous pointer eliminates the need for O(n) extra space to store all values
  - The solution avoids the two-pass approach (traversal + array check) of Solution 1
  - Use double pointer **TreeNode to track the previous node across recursive calls

```go
func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
    return validate(root, &prev)
}

func validate(node *TreeNode, prev **TreeNode) bool {
    if node == nil {
        return true
    }

    if !validate(node.Left, prev) {
        return false
    }

    if *prev != nil && node.Val <= (*prev).Val {
        return false
    }

    *prev = node

    return validate(node.Right, prev)
}
```
