package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	solve1(root)
	for root.Left != nil {
		root = root.Left
	}
	return root
}
