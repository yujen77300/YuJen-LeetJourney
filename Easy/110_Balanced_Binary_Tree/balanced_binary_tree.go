package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


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

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}