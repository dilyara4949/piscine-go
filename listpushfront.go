package piscine

func ListPushFront(l *List, data interface{}) {
	n := NodeL{
		Data: data,
		Next: nil,
	}
	if l.Tail == nil {
		l.Tail = &n
	} else {
		n.Next = l.Head
	}
	l.Head = &n
}
