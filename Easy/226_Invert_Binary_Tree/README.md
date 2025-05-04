# [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)

## Solution: Recursive Approach
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - We visit each node exactly once during the traversal
  - Each node operation (swapping children) takes O(1) time
- **Space Complexity**: O(h) - where h is the height of the tree
  - The space is used by the recursion call stack
  - In worst case (skewed tree), h=n, this could be O(n)
  - In balanced tree, height = log n, this would be O(log n)
- **Approach**:
  1. Base case: If the current node is null, return null
  2. Swap the left and right children of the current node
  3. Recursively invert the left subtree
  4. Recursively invert the right subtree
  5. Return the root node

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return root
	}

  temp := root.Left
  root.Left = root.Right
  root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)

	return root

}

```