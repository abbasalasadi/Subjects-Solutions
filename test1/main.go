package main

import (
	"fmt"
)

type Stack struct{
	Item int
	Next *Stack
}

func (s *Stack) Push(item int) *Stack {
	return &Stack{
		Item: item,
		Next: s,
	}
}

func main() {
	var s *Stack
	for i:=0; i<10; i++{
		s=s.Push(i)
		fmt.Println(s.Item)
	}
}