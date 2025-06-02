// allowedFunctions	(3)[ "os.*", "github.com/01-edu/z01.PrintRune", "--allow-builtin" ]

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func Print(arr []string) {
	for i, c := range arr {
		for _, v := range c {
			z01.PrintRune(v)
		}
		if i != len(arr)-1 && c != "" {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func Split(str string) (arr []string) {
outer:
	for len(str) > 0 {
		for i, c := range str {
			if c == ' ' || c == '\t' {
				if str[:i] != "" {
					arr = append(arr, str[:i])
				}
				str = str[i+1:]
				break
			}
			if i == len(str)-1 {
				arr = append(arr, str[:i+1])
				break outer
			}
		}
	}
	return arr
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		z01.PrintRune('\n')
		return
	}
	str := arg[0]
	Print(Split(str))
}
