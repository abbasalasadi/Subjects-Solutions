// allowedFunctions	(5)[ "strconv.Atoi", "os.*", "fmt.*", "strings.Split", "--allow-builtin" ]

package main

import (
	"fmt"
	"os"
)

func Match(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	}
	return '0'
}

type Node struct {
	bracket rune
	next    *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func Push(l *List, Bracket rune) {
	newBracket := &Node{bracket: Bracket}
	if l.Head == nil {
		l.Head = newBracket
		l.Tail = newBracket
	} else {
		newBracket.next = l.Head
		l.Head = newBracket
	}
}

func Pop(l *List) *Node {
	if l == nil || l.Head == nil {
		return nil
	}
	Bracket := l.Head
	l.Head = l.Head.next
	if l.Head == nil {
		l.Tail = nil
	}
	return Bracket
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

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		return
	}
	l := &List{}
	ref1 := "([{"
	ref2 := ")]}"
	for i := range arg {
		for _, c := range arg[i] {
			for _, v := range ref1 {
				if c == v {
					Push(l, v)
				}
			}
			for _, v := range ref2 {
				if c == v {
					open := Pop(l)
					if open == nil {
						fmt.Println("Error")
						return
					}
					// fmt.Println(string(open.bracket), string(v))
					if Match(open.bracket) != v {
						fmt.Println("Error")
						return
					}
				}
			}
		}
		if ListSize(l) != 0 {
			fmt.Println("Error")
			return
		}
		fmt.Println("OK")
	}
}
