package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	prev := &NodeL{}
	prev = nil
	h := l.Head
	for h != nil {
		if data_ref == h.Data {
			if prev == nil {
				l.Head = l.Head.Next
				// h = h.Next
			} else {
				prev.Next = h.Next
				h = prev

			}
		} else {
			prev = h
		}
		h = h.Next
	}
}
