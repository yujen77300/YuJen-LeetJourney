# [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)

## Solution 1: Level Order Traversal (BFS)
- **Time Complexity**: O(n) - where n is the number of nodes in the tree
  - Visit each node exactly once during the breadth-first traversal
  - Each node is processed in constant time
- **Space Complexity**: O(w) - where w is the maximum width of the tree
  - The queue stores nodes at the current level, which can be at most w nodes
  - In the worst case (complete binary tree), w can be up to n/2
- **Approach**:
  1. Use a queue to perform level-order traversal (BFS)
  2. For each level, process all nodes in the current level
  3. Keep track of the last node processed in each level
  4. The rightmost node in each level is added to the result
  5. Continue until all levels are processed
- **Key Insights**:
  - Level-order traversal naturally groups nodes by their depth
  - The last node processed in each level is the rightmost visible node


```go
func countSubstrings(s string) int {
    answer := 0

    expand := func(l, r int) {
        for l >= 0 && r <= len(s) - 1 && s[l] == s[r] {
            answer ++
            l--
            r ++
        }
    }

    for i := 0 ; i <= len(s)-1 ; i++ {
        expand(i,i)
        expand(i,i+1)
    }

    return answer
}
```


