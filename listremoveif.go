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

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}

	var prev *NodeL
	current := l.Head
	for current != nil {
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}

			if current.Next == nil {
				l.Tail = prev
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
