# [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## Solution 1: Recursive Approach
- **Time Complexity**: O(n²) - where n is the number of nodes in the tree
  - At each recursive call, we need to find the root in the inorder array: O(n) in worst case
  - We need to do this for each of the n nodes in the tree
  - Total time complexity: O(n²)
- **Space Complexity**: O(n)
  - Recursive call stack can go as deep as the height of the tree: O(h), which is O(n) in worst case
  - We also create new slices of preorder and inorder arrays in each recursive call: O(n)
- **Approach**:
  1. Use the first element of preorder array as the root of current subtree
  2. Find this root value's index in the inorder array
  3. Elements to the left of this index in inorder array form the left subtree
  4. Elements to the right of this index in inorder array form the right subtree
  5. Recursively build left and right subtrees using corresponding segments of preorder and inorder arrays
  6. Return the constructed tree

check every single substring

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootValue := preorder[0]
	root := &TreeNode{
		Val: rootValue,
	}

	inorderIndex := indexOf(inorder, rootValue)

	root.Left = buildTree(preorder[1:1+inorderIndex], inorder[:inorderIndex])
	root.Right = buildTree(preorder[1+inorderIndex:], inorder[inorderIndex+1:])

	return root

}

func indexOf(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}
```