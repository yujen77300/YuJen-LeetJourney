package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
