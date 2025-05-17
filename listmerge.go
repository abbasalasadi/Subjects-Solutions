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

func ListMerge(l1 *List, l2 *List) {
	if l2 == nil || l2.Head == nil {
		return
	}

	current := l2.Head
	for current != nil {
		ListPushBack(l1, current.Data)
		current = current.Next
	}
}
