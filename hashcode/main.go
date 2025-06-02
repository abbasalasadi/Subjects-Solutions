package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.HashCode("A"))
	fmt.Println(p.HashCode("AB"))
	fmt.Println(p.HashCode("BAC"))
	fmt.Println(p.HashCode("Hello World"))
}
