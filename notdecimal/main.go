package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Print(p.NotDecimal("0.1"))
	fmt.Print(p.NotDecimal("174.2"))
	fmt.Print(p.NotDecimal("0.1255"))
	fmt.Print(p.NotDecimal("1.20525856"))
	fmt.Print(p.NotDecimal("-0.0f00d00"))
	fmt.Print(p.NotDecimal(""))
	fmt.Print(p.NotDecimal("-19.525856"))
	fmt.Print(p.NotDecimal("-1952"))
}
