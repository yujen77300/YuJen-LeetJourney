package binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		queueLength := len(queue)

		for i := 0; i < queueLength; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if i == queueLength-1 {
				result = append(result, node.Val)
			}
		}
	}

	return result

}
