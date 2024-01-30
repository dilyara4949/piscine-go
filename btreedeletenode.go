package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	solve1(root)
	return delete(root, node)
}

func delete(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data > node.Data {
		root.Left = delete(root.Left, node)
		return root
	} else if root.Data < node.Data {
		root.Right = delete(root.Right, node)
		return root
	}
	if root.Left == nil {
		temp := root.Right
		return temp
	} else if root.Right == nil {
		temp := root.Left
		return temp
	} else {
		succParent := root
		succ := root.Right
		for succ.Left != nil {
			succParent = succ
			succ = succ.Left
		}
		if succParent != root {
			succParent.Left = succ.Right
		} else {
			succParent.Right = succ.Right
		}
		root.Data = succ.Data
		return root
	}
}
