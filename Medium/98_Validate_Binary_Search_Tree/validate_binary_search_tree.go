package validatebinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	nums := []int{}
	inOrderTraversal(root, &nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}

	return true

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
