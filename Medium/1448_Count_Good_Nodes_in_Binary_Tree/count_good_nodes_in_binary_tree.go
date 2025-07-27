package countgoodnodesinbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	var dfs func(node *TreeNode, maxSoFar int) int

	dfs = func(node *TreeNode, maxSoFar int) int {
		res := 0
		if node == nil {
			return 0
		}

		if node.Val >= maxSoFar {
			res = 1
		}

		maxSoFar = max(maxSoFar, node.Val)
		res += dfs(node.Left, maxSoFar)
		res += dfs(node.Right, maxSoFar)

		return res
	}

	return dfs(root, root.Val)

}
