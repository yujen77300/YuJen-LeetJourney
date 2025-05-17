# [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversalg/)

## Solution 1: BFS
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Each node is processed exactly once when it's dequeued
  - Adding and removing elements from the queue takes O(1) time
- **Space Complexity**: O(n) - for storing the queue and result
  - In the worst case (a complete binary tree's last level), the queue can hold up to n/2 nodes
  - The result list stores all n node values
- **Approach**:
  1. Use a queue data structure (FIFO - First In, First Out)
  2. Start by adding the root node to the queue
  3. Process nodes level by level, tracking the number of nodes at each level
  4. For each node, add its value to the current level list and enqueue its children
  5. After processing all nodes at a level, add the level list to the result
  6. Continue until the queue is empty


```go
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		queueLength := len(queue)
		level := []int{}

		for i := 0; i < queueLength; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
```