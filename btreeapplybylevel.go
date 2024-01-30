package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelCount(root)
	for i := 0; i < h; i++ {
		solve(root, i, f)
	}
}

func solve(root *TreeNode, i int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 0 {
		f(root.Data)
		return
	}
	solve(root.Left, i-1, f)
	solve(root.Right, i-1, f)
}
