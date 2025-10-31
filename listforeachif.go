package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	for node := l.Head; node != nil; node = node.Next {
		if cond(node) {
			f(node)
		}
	}
}

func IsAlNode(node *NodeL) bool {
	str, ok := node.Data.(string)
	if !ok || len(str) == 0 {
		return false
	}
	for i := 0; i < len(str); i++ {
		c := str[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			return true
		}
	}
	return false
}
