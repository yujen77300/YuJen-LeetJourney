# [Count Good Nodes in Binary Tree](https://leetcode.com/problems/count-good-nodes-in-binary-tree/)

## Solution 1: Recursive DFS Approach
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Visit each node exactly once during the depth-first traversal
  - Each node is processed in constant time
- **Space Complexity**: O(h) - where h is the height of the tree
  - Recursion stack depth equals the height of the tree
  - In worst case (skewed tree), h can be n; in best case (balanced tree), h is log(n)
- **Approach**:
  1. Use DFS to traverse the tree while maintaining the maximum value seen so far on the path
  2. For each node, check if its value is greater than or equal to the maximum value on the path
  3. If yes, increment the count of good nodes
  4. Update the maximum value for the current path and recursively traverse left and right subtrees
  5. Return the total count of good nodes
- **Key Insights**:
  - Track the maximum value encountered on the path from root to current node


```go
func goodNodes(root *TreeNode) int {

    var dfs func(node *TreeNode, maxSoFar int) int

    dfs = func(node *TreeNode, maxSoFar int)int{
        res := 0
        if node == nil {
            return 0
        }

        if node.Val >= maxSoFar {
            res = 1
        }

        maxSoFar = max(maxSoFar, node.Val)
        // Recursively count good nodes in left and right subtrees
        res += dfs(node.Left, maxSoFar)
        res += dfs(node.Right, maxSoFar)

        return res
    }

    return dfs(root, root.Val)

}

```




## Solution 2: Iterative Stack Approach
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Visit each node exactly once during the iterative traversal
  - Each node is processed in constant time
- **Space Complexity**: O(h) - where h is the height of the tree
  - Stack stores nodes along with their maximum path values
  - In worst case (skewed tree), stack size can be n; in best case (balanced tree), it's log(n)
- **Approach**:
  1. Use an explicit stack to simulate DFS traversal
  2. Store both the tree node and the maximum value seen so far in the stack
  3. For each node popped from stack, check if it's a good node
  4. Update the maximum value and push children to stack
  5. Continue until stack is empty
- **Key Insights**:
  - Recursive DFS is concise, but on very deep trees it might cause a stack overflow. Using an iterative DFS with an explicit stack is a safer alternative in such cases.
  - Custom stack node structure stores both tree node and path maximum

```go
func goodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    type StackNode struct {
        node     *TreeNode
        maxSoFar int
    }

    stack := []StackNode{{root, root.Val}}
    goodNodesCount := 0

    for len(stack) > 0 {
        // Pop
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        node := curr.node
        maxSoFar := curr.maxSoFar

        if node.Val >= maxSoFar {
            goodNodesCount++
        }

        newMax := max(maxSoFar, node.Val)

        if node.Right != nil {
            stack = append(stack, StackNode{node.Right, newMax})
        }
        if node.Left != nil {
            stack = append(stack, StackNode{node.Left, newMax})
        }
    }

    return goodNodesCount
}
```


