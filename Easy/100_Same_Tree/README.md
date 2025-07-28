# [Same Tree](https://leetcode.com/problems/same-tree/)

## Solution: Recursive Approach
- **Time Complexity**: O(n) - where n is the total number of nodes in both trees
  - We need to visit each node in both trees exactly once during the traversal
- **Space Complexity**: O(h) - where h is the height of the tree 
  - In the worst case (skewed tree), this could be O(n)
  - The space is used by the recursion call stack
- **Approach**:
  1. Base cases:
     - If both trees are nil, return true
     - If only one tree is nil, return false
  2. Check if the current node values are equal
  3. Recursively check if the left subtrees are the same
  4. Recursively check if the right subtrees are the same
  5. Return true only if all checks pass

### Go
```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
```

### Python
```python
class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if p is None and q is None :
            return True
        if p is None or q is None :
            return False
        if p.val != q.val :
            return False

        return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
```

**Note**: In Python, use is None instead of == None when checking for None values. The is operator checks for object identity (whether two variables refer to the same object), while == checks for equality. Since None is a singleton in Python, `is None` is more efficient and is considered the Pythonic way.