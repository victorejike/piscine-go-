package piscine

func ListSize(l *List) int {
	count := 0
	for node := l.Head; node != nil; node = node.Next {
		count++
	}
	return count
}
