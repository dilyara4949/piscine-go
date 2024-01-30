package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	solve1(root)
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func solve1(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Right != nil {
		root.Right.Parent = root
		solve1(root.Right)
	}
	if root.Left != nil {
		root.Left.Parent = root
		solve1(root.Left)
	}
}
