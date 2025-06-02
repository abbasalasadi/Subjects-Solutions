// allowedFunctions	(3)[ "github.com/01-edu/z01.PrintRune", "unicode.IsGraphic", "--allow-builtin" ]

package piscine

import (
	"unicode"

	"github.com/01-edu/z01"
)

func Print(arr []byte) {
	for _, c := range arr {
		if unicode.IsGraphic(rune(c)) {
			z01.PrintRune(rune(c))
		} else {
			z01.PrintRune('.')
		}
	}
}

func PrintHex(hex []string) {
	length := 0
	for len(hex) > 0 {
		if len(hex) > 4 {
			length = 4
		} else {
			length = len(hex)
		}
		for i := 0; i < length; i++ {
			if hex[i] != "" {
				Print([]byte(hex[i]))
			} else {
				Print([]byte{48, 48})
			}
			if i != length {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
		z01.PrintRune('\n')
		if len(hex) > 4 {
			hex = hex[4:]
		} else {
			break
		}
	}
}

func ToHex(arr []byte) (hex []string) {
	ref := "0123456789abcdef"
	for _, c := range arr {
		n := int(c)
		s := ""
		for n > 0 {
			rem := n % 16
			s = string(ref[rem]) + s
			n /= 16
		}
		hex = append(hex, s)
	}
	return hex
}

func PrintMemory(arr [10]byte) {
	arr1 := arr[:]
	hex := ToHex(arr1)
	PrintHex(hex)
	Print(arr1)
}
