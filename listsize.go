package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	NewNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = NewNode
		l.Tail = NewNode
	} else {
		NewNode.Next = l.Head
		l.Head = NewNode
	}
}

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	count := 0
	current := l.Head
	for current != nil {
		current = current.Next
		count++
	}
	return count
}
