// allowedFunctions	(3)[ "os.Args", "github.com/01-edu/z01.PrintRune", "--allow-builtin" ]

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func Trim(s string) string {
	var str string
	for i, c := range s {
		if c != ' ' {
			str = s[i:]
			break
		}
	}
	for i, c := range str {
		if c == ' ' {
			str = str[:i]
			break
		}
	}
	return str
}

func Slice(s string) (list []string) {
Outer:
	for len(s) > 0 {
		for i, v := range s {
			if v == ' ' {
				word := Trim(s[:i])
				if len(word) != 0 {
					list = append(list, Trim(s[:i]))
				}
				s = s[i+1:]
				break
			}
			if i == len(s)-1 {
				list = append(list, Trim(s))
				break Outer
			}
		}
	}
	return list
}

func Print(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		z01.PrintRune('\n')
		return
	}
	str := arg[0]
	Slice := Slice(str)
	Slice = append(Slice[1:], Slice[0])

	str = ""
	for i, v := range Slice {
		str = str + v
		if i != len(Slice)-1 {
			str = str + " "
		}
	}
	Print(str)
}
