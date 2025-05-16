package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func Enqueue(l *List, data int) {
	NewNode := &Node{Value: data}
	if l.Tail == nil {
		l.Head = NewNode
		l.Tail = NewNode
	} else {
		l.Tail.Next = NewNode
		l.Tail = NewNode
	}
}

func Push(l *List, data int) {
	NewNode := &Node{Value: data}
	if l.Head == nil {
		l.Head = NewNode
		l.Tail = NewNode
	} else {
		NewNode.Next = l.Head
		l.Head = NewNode
	}
}

func Dequeue(l *List) int {
	if l.Head == nil {
		return -1
	}
	data := l.Head.Value
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	return data
}

func Size(l *List) int {
	count := 0
	current := l.Head
	for current != nil {
		current = current.Next
		count++
	}
	return count
}

func Clear(l *List) {
	if l == nil {
		return
	}
	l.Head = nil
	l.Tail = nil
}

func Find(l *List, ref int) int {
	if l == nil {
		return -1
	}
	current := l.Head
	count := 0
	for current != nil {
		if current.Value == ref {
			return count
		}
		current =current.Next
		count++
	}
	return count
}

func main() {
	link := &List{}

	i := 0
	for i < 10 {
		Push(link, i)
		i++
	}

	// Clear(link)

	fmt.Println("size: ", Size(link))
	fmt.Println("item dequeued: ", Dequeue(link))
	fmt.Println("item dequeued: ", Dequeue(link))
	fmt.Println("size: ", Size(link))

	fmt.Println(Find(link,0))
}
