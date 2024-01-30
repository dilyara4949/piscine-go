package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	} else if elem > root.Data {
		if root.Right != nil {
			root.Right.Parent = root
		}
		return BTreeSearchItem(root.Right, elem)
	} else {
		if root.Left != nil {
			root.Left.Parent = root
		}
		return BTreeSearchItem(root.Left, elem)
	}
}
