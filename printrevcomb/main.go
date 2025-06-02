// allowedFunctions	[ "github.com/01-edu/z01.PrintRune", "--no-lit=\\d{3}," ]

package main

import "github.com/01-edu/z01"

func Print(nums ...int) {
	for _, c := range nums {
		z01.PrintRune(rune(c) + '0')
	}
}

func main() {
	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			for k := 9; k >= 0; k-- {
				if i > j && j > k {
					Print(i, j, k)
					if !(i == 2 && j == 1 && k == 0) {
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
