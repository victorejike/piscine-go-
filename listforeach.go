package piscine

func ListForEach(l *List, f func(*NodeL)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node)
	}
}

func Add2_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v + 2
	case string:
		node.Data = v + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch v := node.Data.(type) {
	case int:
		node.Data = v - 3
	case string:
		if len(v) > 0 {
			last := v[len(v)-1]
			if last == '2' || last == '3' {
				node.Data = v[:len(v)-1]
				return
			}
		}
		node.Data = v
	}
}
