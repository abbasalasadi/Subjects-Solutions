// allowedFunctions	(3)[ "os.Args", "github.com/01-edu/z01.PrintRune", "--allow-builtin" ]

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func Print(list []string) {
	for i := len(list) - 1; i >= 0; i-- {
		for _, v := range list[i] {
			z01.PrintRune(v)
		}
		if i != 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}
	str := arg[0]

	list := []string{}
	for len(str) > 0 {
		for i, v := range str {
			if v == ' ' || i == len(str)-1 {
				list = append(list, str[:i])
				str = str[i+1:]
				break
			}
		}
	}
	Print(list)
}
