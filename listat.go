package piscine

// import "fmt"

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

func ListSize(l *List) int {
	if l.Head == nil {
		return -1
	}
	count := 0
	current := l.Head
	for l.Head != nil {
		current = current.Next
		count++
	}
	return count
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	current := l
	count := 0
	for count != pos {
		if current == nil {
			return nil
		}
		count++
		current = current.Next
	}
	return current
}
