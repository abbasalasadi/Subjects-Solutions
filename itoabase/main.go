package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ItoaBase(10, 2)) // 1010
	fmt.Println(piscine.ItoaBase(255, 16)) // FF
	fmt.Println(piscine.ItoaBase(-42, 4)) // -222
	fmt.Println(piscine.ItoaBase(123, 10)) // 123
	fmt.Println(piscine.ItoaBase(0, 8)) // 0
	fmt.Println(piscine.ItoaBase(255, 2)) // 11111111
	fmt.Println(piscine.ItoaBase(-255, 16)) // -FF
	fmt.Println(piscine.ItoaBase(15, 16)) // F
	fmt.Println(piscine.ItoaBase(10, 4)) // 22
	fmt.Println(piscine.ItoaBase(255, 10)) // 255
}
