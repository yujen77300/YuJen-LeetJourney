# [Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree//)

## Solution: Recursive Approach
- **Time Complexity**: O(m * n) - where m is the number of nodes in the root tree and n is the number of nodes in the subRoot tree
  - In worst case, we might need to check each node in the root tree as a potential match for the subRoot tree
  - For each potential match, we perform a tree comparison that takes O(n) time
- **Space Complexity**: O(h) - where h is the height of the root tree
  - The space is used by the recursion call stack
  - In worst case (skewed tree), this could be O(m)
- **Approach**:
  1. Base cases:
     - If subRoot is null, any tree contains an empty subtree, return true
     - If root is null but subRoot is not, return false (can't find subtree in empty tree)
  2. For each node in the main tree:
     - Check if the current node and its subtree match the subRoot tree
     - If match found, return true
     - Otherwise, recursively check left and right subtrees
  3. Return true if a match is found in any subtree
- **Key Insights**:
  - The problem is solved by combining two tree operations:
    - Tree traversal to find potential matching subtrees
    - Tree comparison to verify if two trees are identical

```go
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if sameTree(root, subRoot) {
		return true
	}

	return (isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot))
}

func sameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root != nil && subRoot != nil && root.Val == subRoot.Val {
		return (sameTree(root.Left, subRoot.Left) && sameTree(root.Right, subRoot.Right))
	}

	return false
}

```