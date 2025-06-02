package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.WeAreUnique("foo", "boo"))
	fmt.Println(p.WeAreUnique("", ""))
	fmt.Println(p.WeAreUnique("abc", "def"))
}
