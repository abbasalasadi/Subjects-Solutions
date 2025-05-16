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

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

func ListForEach(l *List, f func(*NodeL)) {
	if l == nil || l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}
