# [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)

## Solution: Recursive Approach
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - We need to visit each node exactly once during the traversal
- **Space Complexity**: O(h) - where h is the height of the tree
  - In the worst case (skewed tree), this could be O(n)
  - The space is used by the recursion call stack
- **Approach**:
  1. Use a helper function to track the current depth during traversal
  2. Base case: If the current node is null, return the current depth
  3. Recursively calculate the depth of both left and right subtrees
  4. Return the maximum of the two depths
  5. The final result represents the maximum depth of the tree


```go
func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return max(helper(root.Left, depth+1), helper(root.Right, depth+1))
}
```

Without helper function
```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
```

```python
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None :
            return 0
        return max(self.maxDepth(root.left), self.maxDepth(root.right)) + 1
````