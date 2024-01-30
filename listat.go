package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for l != nil && pos > 0 {
		l = l.Next
		pos--
	}
	return l
}
