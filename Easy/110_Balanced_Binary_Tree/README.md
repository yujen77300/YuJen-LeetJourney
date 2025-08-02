# [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)

## Solution 1: Bottom up DFS
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Each node is visited exactly once during the post-order traversal
  - Early termination prevents unnecessary computation once imbalance is detected
- **Space Complexity**: O(h) - where h is the height of the tree
  - Recursion call stack space, which is O(log n) for balanced trees and O(n) for worst case (skewed tree)
  - No additional data structures needed
- **Approach**:
  1. Use post-order traversal to calculate subtree heights
  2. For each node, recursively calculate the height of left and right subtrees
  3. Check if the absolute difference between left and right heights exceeds 1
  4. If imbalance is found at any node, mark the tree as unbalanced and stop further computation
  5. Return the height of current subtree for parent node calculations
- **Key Insights**:
  - Bottom-up approach allows us to detect imbalance early and avoid redundant calculations
  - Post-order traversal ensures we have height information from children before processing parent
  - Custom abs function avoids float conversion issues with math.Abs

```go
func isBalanced(root *TreeNode) bool {
    balanced := true

    var height func(node *TreeNode) int
    height = func(node *TreeNode) int {
        if !balanced {
            return 0 
        }

        if node == nil {
            return 0
        }

        leftHeight := height(node.Left)
        rightHeight := height(node.Right)

        if abs(leftHeight - rightHeight) > 1 {
            balanced = false
        }

        return 1 + max(leftHeight, rightHeight)
    }

    height(root)
    return balanced
}
```

比較高度 用post order

只要設計「遞迴函式的回傳值」就是「子樹高度」，
就可以在 DFS 的過程中：


| 類型          | 處理順序                   | 代表例子                                   |
| ------------- | -------------------------- | ------------------------------------------ |
| Top-down      | 根 → 左 → 右（pre-order）  | 常用於剪枝、傳遞參數往下帶（例如深度）     |
| **Bottom-up** | 左 → 右 → 根（post-order） | ✅ 適合處理子問題回傳給上層（如高度、平衡） |


math.Abs是float要小心