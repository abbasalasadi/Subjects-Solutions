// allowedFunctions	(3)[ "github.com/01-edu/z01.PrintRune", "os.Args", "--allow-builtin" ]

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}
	str := arg[0]
	find := rune(arg[1][0])
	replace := rune(arg[2][0])

	result := []rune{}
	for _, c := range str {
		if c == find {
			result = append(result, replace)
		} else {
			result = append(result, c)
		}
	}
	Print(string(result))
}
