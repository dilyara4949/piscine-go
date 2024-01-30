package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	x := BTreeLevelCount(root.Left)
	y := BTreeLevelCount(root.Right)
	if x > y {
		return x + 1
	}
	return y + 1
}
