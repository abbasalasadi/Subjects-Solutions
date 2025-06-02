package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.DigitLen(100, 10))
	fmt.Println(p.DigitLen(100, 2))
	fmt.Println(p.DigitLen(-100, 16))
	fmt.Println(p.DigitLen(100, -1))
}
