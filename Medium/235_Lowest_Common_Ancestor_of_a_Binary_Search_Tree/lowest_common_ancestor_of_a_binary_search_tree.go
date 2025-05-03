package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pointer := root

	for pointer != nil {
		if p.Val > pointer.Val && q.Val > pointer.Val {
			pointer = pointer.Right
		} else if p.Val < pointer.Val && q.Val < pointer.Val {
			pointer = pointer.Left
		} else {
			return pointer
		}
	}

	return pointer
}
