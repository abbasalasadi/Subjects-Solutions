//allowedFunctions	(3)[ "os.Args", "fmt.Printf", "--allow-builtin" ]

package main

import (
	"fmt"
	"os"
)

func Valid(s string) (num int, err bool) {
	if s == "0" {
		return 0, false
	}
	for _, v := range s {
		if v < '0' || v > '9' {
			return 0, false
		}
		num = num*10 + int(v-'0')
	}
	if num >= 4000 {
		return 0, false
	}
	return num, true
}

type Node struct {
	num  int
	next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func Enqueue(l *List, n int) {
	newNode := &Node{num: n}
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.next = newNode
		l.Tail = newNode
	}
}

func Dequeue(l *List) (n int) {
	if l == nil || l.Head == nil {
		return -1
	}
	n = l.Head.num
	l.Head = l.Head.next

	if l.Head == nil {
		l.Tail = nil
	}
	return n
}

func Cast(n, dec int) (Rom, Text string) {
	Ref := []string{"I", "V", "X", "L", "C", "D", "M"}
	switch {
	case n >= 1 && n <= 3:
		for i := 0; i < n; i++ {
			Rom = Rom + Ref[0+2*dec]
			Text = ""
			for i, v := range Rom {
				Text = Text + string(v)
				if i != len(Rom)-1 {
					Text = Text + "+"
				}
			}
		}
	case n == 4:
		Rom = Rom + Ref[0+2*dec] + Ref[1+2*dec]
		Text = "(" + string(Rom[1]) + "-" + string(Rom[0]) + ")"
	case n == 5:
		Rom = Rom + Ref[1+2*dec]
		Text = Rom
	case n >= 6 && n <= 8:
		Rom = Rom + Ref[1+2*dec]
		for i := 0; i < n-5; i++ {
			Rom = Rom + Ref[0+2*dec]
		}
		Text = ""
		for i, v := range Rom {
			Text = Text + string(v)
			if i != len(Rom)-1 {
				Text = Text + "+"
			}
		}
	case n == 9:
		Rom = Rom + Ref[0+2*dec] + Ref[2+2*dec]
		Text = "(" + string(Rom[1]) + "-" + string(Rom[0]) + ")"
	case n == 10:
		Rom = Rom + Ref[2+2*dec]
		Text = Rom
	}
	return Rom, Text
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Printf("Invalid No of arguments\n")
	}

	num, vld := Valid(arg[0])
	if !vld {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}

	link := &List{}
	for num > 0 {
		Enqueue(link, num%10)
		num /= 10
	}

	i := 0
	Rom := ""
	Text := []string{}
	for link.Head != nil {
		R, T := Cast(Dequeue(link), i)
		Rom = R + Rom
		Text = append([]string{T},Text...)
		i++
	}
	for i,v:=range Text{
		fmt.Printf("%v", v)
		if i != len(Text)-1 {
			fmt.Printf("+")
		} else {
			fmt.Printf("\n")
		}
	}
	
	fmt.Printf("%v", Rom)
}

