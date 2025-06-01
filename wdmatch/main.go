package main

import (
	"fmt"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) <= 1 {
		return
	}
	one := arg[0]
	two := arg[1]
	check := two
	found := false
	stop := false
	for i, c := range one {
		for j, v := range check {
			fmt.Println("j: ", j)
			if len(one) > 0 && len(two) == 0 {
				stop = true
				fmt.Println(1)
				break
			} else if i == len(one)-1 && c == v {
				found = true
				stop = true
				fmt.Println(2)
				break
			} else if c == v && j+1 < len(check)-1 {
				check = check[j+1:]
				fmt.Println(3)
				break
			}
		}
		fmt.Println(string(c))
		fmt.Println("len(check)", len(check))
		fmt.Println(check)
		if stop {
			break
		}
	}
	if found {
		fmt.Println("len(two)", len(two))
		for _, c := range one {
			z01.PrintRune(c)
		}
	}
}
