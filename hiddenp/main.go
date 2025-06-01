// os.Args", "github.com/01-edu/z01.PrintRune", "--allow-builtin"
package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	// If the number of arguments is different from 2, the program should display nothing.
	if len(arg) != 2 {
		return
	}

	// If s1 is an empty string, it is considered hidden in any string.
	if arg[0] == "" {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	// If s1 is hidden in s2, the program should display 1 followed by a newline.
	s1 := arg[0]
	s2 := arg[1]
	temp := s2
	found := false
	stop := false
	// fmt.Println("len(s1)-1: ", len(s1)-1)

	for i, c := range s1 {
		// stop function
		if i == len(s1)-1 || len(temp) == 0 { //last item in s1
			for j := 0; j < len(temp); j++ {
				// fmt.Printf("loop 1, i: %v, j: %v, c: %v, temp[j]: %v , len(tem)-1: %v \n", i, j, string(c), string(rune(temp[j])), len(temp)-1)
				// fmt.Println(temp)
				if c == rune(temp[j]) {
					found = true
					break
				}
				if j == len(temp)-1 {
					stop = true
					break
				}
			}
		} else { // not the last item in s1
			for j := 0; j < len(temp); j++ {
				// fmt.Printf("loop 2, i: %v, j: %v, c: %v, temp[j]: %v , len(tem)-1: %v \n", i, j, string(c), string(rune(temp[j])), len(temp)-1)
				// fmt.Println(temp)
				if j == len(temp)-1 {
					stop = true
					break
				}
				if c == rune(temp[j]) {
					temp = temp[j+1:]
					break
				}

				if j == len(temp)-1 {
					stop = true
					break
				}
			}
		}
		if stop {
			// fmt.Println("not found")
			z01.PrintRune('0')
			break
		}
		if found {
			// fmt.Println("found")
			z01.PrintRune('1')
			break
		}
	}
	// fmt.Println("found: ", found)
	// fmt.Println("stop: ", stop)

}
