package maximumdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return max(helper(root.Left, depth+1), helper(root.Right, depth+1))
}
