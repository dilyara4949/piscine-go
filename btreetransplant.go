package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	repl(root, node, rplc)
	if node == root {
		root = node
	}
	solve1(root)
	return root
}

func repl(root, node, rplc *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left == node {
		root.Left = rplc
		return
	} else if root.Right != nil && root.Right == node {
		root.Right = rplc
		return
	}
	repl(root.Left, node, rplc)
	repl(root.Right, node, rplc)
}
