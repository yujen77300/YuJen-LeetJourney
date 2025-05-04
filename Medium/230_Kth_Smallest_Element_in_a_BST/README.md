# [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/)

## Solution 1: In-order Traversal with Array
- **Time Complexity**: O(n) - where n is the number of nodes in the BST
  - We need to visit each node exactly once during the in-order traversal
  - Accessing the kth element in the array is O(1)
- **Space Complexity**: O(n) - extra space to store all nodes' values
  - We store all n values from the BST in an array
- **Approach**:
  1. Perform an in-order traversal of the BST, storing all values in an array
  2. In-order traversal of a BST visits nodes in ascending order
  3. Return the kth element (index k-1) from the sorted array
- **Key Insights**:
  - In-order traversal of a BST naturally produces elements in sorted order
  - Simple approach but uses extra space to store all values

```go
func kthSmallest(root *TreeNode, k int) int {
    nums := []int{}
    inOrderTraversal(root, &nums)

    return nums[k-1]

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




## Solution 2: In-order Traversal with Closure
- **Time Complexity**: O(h + k) - where h is the height of the tree and k is the input parameter
  - In worst case, we need to traverse down to the leftmost node (O(h))
  - Then we visit at most k nodes to find the kth smallest
  - In a balanced BST, this is O(log n + k)
- **Space Complexity**: O(h) - where h is the height of the tree
  - The space is used by the recursion call stack
  - In worst case (skewed tree), this could be O(n)
- **Approach**:
  1. Perform an in-order traversal of the BST
  2. Keep track of the number of nodes visited
  3. When we've visited exactly k nodes, we've found the kth smallest element
  4. Stop traversal early once we find the target
- **Key Insights**:
  - Optimizes the first approach by avoiding storing all node values
  - Uses closure to maintain state across recursive calls
  - Early stopping saves time when k is much smaller than n


```go
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0

	var inOrderTraversal func(node *TreeNode)
	inOrderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		inOrderTraversal(node.Left)

		count++
		if count == k {
			result = node.Val
			return
		}

		inOrderTraversal(node.Right)
	}

	inOrderTraversal(root)

	return result
}

```


