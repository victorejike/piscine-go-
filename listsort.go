package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	for i := l; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}
	return jn
}
