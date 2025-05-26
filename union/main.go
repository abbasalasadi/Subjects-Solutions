// allowedFunctions	(3)[ "os.Args", "github.com/01-edu/z01.PrintRune", "--allow-builtin" ]

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		z01.PrintRune('\n')
		return
	}

	a := arg[0]
	b := arg[1]

	str := a + b

	m := make(map[rune]bool)
	for _, v := range str {
		if _, exist := m[v]; !exist {
			m[v] = true
			z01.PrintRune(v)
		}
	}
}
