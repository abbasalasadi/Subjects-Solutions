// allowedFunctions	(3)[ "os.Args", "github.com/01-edu/z01.PrintRune", "--allow-builtin" ]

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}
	str := arg[0]
	vowels := "aeiou"
	Index := -1
outer:
	for _, c := range vowels {
		for i, v := range str {
			if c == v {
				Index = i
				break outer
			}
		}
	}
	Txt := ""
	switch {
	case Index == -1:
		Print("No vowels")
	case Index == 0:
		Txt = str + "ay"
		Print(Txt)
	case Index > 0:
		Txt = str[Index:] + str[:Index] + "ay"
		Print(Txt)
	}

}
