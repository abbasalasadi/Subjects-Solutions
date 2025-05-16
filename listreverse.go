package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushBack(l *List, data interface{}) {
// 	NewNode := &NodeL{Data: data}
// 	if l.Tail == nil {
// 		l.Head = NewNode
// 		l.Tail = NewNode
// 	} else {
// 		l.Tail.Next = NewNode
// 		l.Tail = NewNode
// 	}
// }

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var prev *NodeL
	current := l.Head
	l.Tail = l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
