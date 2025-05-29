// allowedFunctions	(5)[ "strconv.Atoi", "os.*", "fmt.*", "strings.Split", "--allow-builtin" ]

package main

import (
	"fmt"
	"os"
)

type Node struct {
	num  int
	next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func Push(l *List, data int) {
	newNode := &Node{num: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.next = l.Head
		l.Head = newNode
	}
}

func Enqueue(l *List, data int) {
	newNode := &Node{num: data}
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.next = newNode
		l.Tail = newNode
	}
}

func Dequeue(l *List) *Node {
	if l == nil || l.Head == nil {
		return nil
	}
	num := l.Head
	l.Head = l.Head.next
	if l.Head == nil {
		l.Tail = nil
	}
	return num
}

func ListSize(l *List) int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func calc(a, b int, opr rune) (result int) {
	switch opr {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	case '%':
		result = a % b
	}
	return result
}

func Valid(s string) bool {
	num := 0
	opr := 0
	for _, v := range s {
		if !((v >= '0' && v <= '9') || v == '+' || v == '-' || v == '*' || v == '/' || v == '%' || v == ' ' || v == '\t') {
			return false
		}
		if v >= '0' && v <= '9' {
			num++
		}
		for _, c := range "+-*/%" {
			if v == c {
				opr++
				break
			}
		}
	}
	return num == opr+1
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("Error")
		return
	}
	if !Valid(arg[0]) {
		fmt.Println("Error")
		return
	}
	l := &List{}
	ref := "+-*/%"
// outer:
	for _, c := range arg[0] {
		if c >= '0' && c <= '9' {
			Enqueue(l, int(c-'0'))
		} else {
			for _, v := range ref {
				if c == v {
					a := Dequeue(l)
					b := Dequeue(l)
					if a == nil || b == nil{
						fmt.Println("Error")
						return
					}
					result := calc(a.num, b.num, c)
					Push(l, result)
					// if ListSize(l) == 1 {
					// 	break outer
					// }
				}
			}
		}
	}
	// fmt.Println("ListSize: ",ListSize(l))
	fmt.Println(Dequeue(l).num)
}

