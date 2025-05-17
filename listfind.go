package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	NewNode := &NodeL{Data: data}
	if l.Tail == nil {
		l.Head = NewNode
		l.Tail = NewNode
	} else {
		l.Tail.Next = NewNode
		l.Tail = NewNode
	}
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l == nil || l.Head == nil {
		return nil
	}
	current := l.Head
	for current != nil {
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}
	return nil
}
