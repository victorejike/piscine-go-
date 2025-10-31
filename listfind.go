package piscine

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for node := l.Head; node != nil; node = node.Next {
		if comp(node.Data, ref) {
			return &node.Data
		}
	}
	return nil
}

func CompStr(a, b interface{}) bool {
	return a == b
}
