package kthsmallestelementinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
