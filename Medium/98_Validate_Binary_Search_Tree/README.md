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
