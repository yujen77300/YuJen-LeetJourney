# [Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

## Solution: Iterative Approach
- **Time Complexity**: O(h) - where h is the height of the BST
  - In the worst case (skewed tree), this could be O(n)
  - We only need to traverse down at most one path from root to leaf
- **Space Complexity**: O(1) - constant extra space
  - We only use a single pointer variable regardless of input size
  - No recursion stack or additional data structures used
- **Approach**:
  1. Start from the root of the BST
  2. Use the BST property to navigate:
     - If both p and q are greater than current node, move to right subtree
     - If both p and q are less than current node, move to left subtree
     - Otherwise, we've found the split point (LCA)
  3. Return the current node when we find the split point
- **Key Insights**:
  - Leverage the BST property: all nodes in left subtree < parent < all nodes in right subtree
  - The LCA is the first node where p and q would go in different directions
  - No need to search the entire tree or track paths - we can navigate directly to the LCA

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pointer := root

	for pointer != nil {
		if p.Val > pointer.Val && q.Val > pointer.Val {
			pointer = pointer.Right
		} else if p.Val < pointer.Val && q.Val < pointer.Val {
			pointer = pointer.Left
		} else {
			return pointer
		}
	}

	return pointer
}
```
