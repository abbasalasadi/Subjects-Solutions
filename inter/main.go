// Allowed functions:  "github.com/01-edu/z01.PrintRune", "os.Args", "--allow-builtin"

package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]

	// If the number of arguments is different from 2, the program displays nothing.
	if len(arg) != 2 {
		return
	}

	// Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.
	// * The display will be followed by a newline ('\n').

	temps1 := arg[0]
	s2 := arg[1]

	m := make(map[rune]int)
	s1 := ""
	for _, c := range temps1 {
		if _, exist := m[c]; exist {
			m[c]++
		} else {
			m[c] = 1
			s1 += string(c)
		}
	}
	// fmt.Println(s1)

	// finding shortest string
	length := min(len(s1), len(s2))
	temp:=""
	for i := 0; i <= length-1; i++ {
		for _,c:=range s2{
			if rune(s1[i])==c {
				temp=temp+string(c)
				break
			}
		}
	}
	fmt.Println(temp)

}
