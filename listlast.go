package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	return node.Data
}
