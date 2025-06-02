package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println(p.IsCapitalized("Hello! How are you?"))
	fmt.Println(p.IsCapitalized("Hello How Are You"))
	fmt.Println(p.IsCapitalized("Whats 4this 100K?"))
	fmt.Println(p.IsCapitalized("Whatsthis4"))
	fmt.Println(p.IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(p.IsCapitalized(""))
}