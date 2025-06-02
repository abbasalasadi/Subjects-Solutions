package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Print(p.ThirdTimeIsACharm("123456789"))
	fmt.Print(p.ThirdTimeIsACharm(""))
	fmt.Print(p.ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(p.ThirdTimeIsACharm("12"))
}
