package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Print(p.PrintIfNot("abcdefz"))
	fmt.Print(p.PrintIfNot("abc"))
	fmt.Print(p.PrintIfNot(""))
	fmt.Print(p.PrintIfNot("14"))
}
