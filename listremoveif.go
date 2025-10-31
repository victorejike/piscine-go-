package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
		return
	}
	prev := l.Head
	curr := l.Head.Next
	for curr != nil {
		if curr.Data == data_ref {
			prev.Next = curr.Next
			if curr == l.Tail {
				l.Tail = prev
			}
		} else {
			prev = curr
		}
		curr = curr.Next
	}
}
