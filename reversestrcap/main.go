package main

import (
	"fmt"
	"os"
)

// func main() {
// 	arg := os.Args[1:]
// 	S := [][]rune{}
// 	for i, c := range arg {
// 		S = append(S, []rune(c))
// 		for j, v := range c {
// 			if v >= 'A' && v <= 'Z' {
// 				S[i][j] = v + 32
// 			}
// 		}
// 	}

// 	for i, c := range S {
// 		for j, v := range S[i] {
// 			if v == ' ' && j != 0 && (v >= 'a' && v <= 'z') {
// 				S[i][j-1] = S[i][j-1] - 32
// 			} else if j == len(c)-1 && j != 0 && (v >= 'a' && v <= 'z') {
// 				S[i][j] = S[i][j] - 32
// 			}
// 		}
// 	}
// 	for i := range S {
// 		for j := range S[i] {
// 			z01.PrintRune(S[i][j])
// 		}
// 		if i != len(S) {
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

func isLetter(c rune) (rune, bool) {
	if c >= 'a' && c <= 'z' {
		return c, true
	}
	if c >= 'A' && c <= 'Z' {
		return c + 32, true
	}
	return c, false
}

func main() {
	arg := os.Args[1:]

	for _, c := range arg {
		temp := ""
		for i, v := range c {
			s, t := isLetter(rune(v))
			if i == len(c)-1 {
				if t {
					// fmt.Println("1: ",temp)
					temp = temp[:i] + string(s-32)
					// fmt.Println("2: ",temp,)
				} else if v == ' ' {
					s, t := isLetter(rune(c[i-1]))
					// fmt.Println("i: ", i)
					if t {
						// fmt.Println("1: ", temp)
						temp = temp[:i-1] + string(s-32) + " "
						// fmt.Println("2: ", temp)
					} else {
						temp = temp[:i-1] + string(s) + " "
					}
				} else {
					temp = temp[:i] + string(s)
				}
			} else {
				if v == ' ' && i != 0 {
					s, t := isLetter(rune(c[i-1]))
					// fmt.Println("i: ", i)
					if t {
						// fmt.Println("1: ", temp)
						temp = temp[:i-1] + string(s-32) + " "
						// fmt.Println("2: ", temp)
					} else {
						temp = temp[:i-1] + string(s) + " "
					}
				} else {
					temp = temp + string(s)
				}
			}

		}
		fmt.Println(temp)
	}
}

// func main() {
// 	arg := os.Args[1:]
// 	// var arr []string
// 	var temp []rune
// 	str := ""
// 	for _, c := range arg {
// 		for _, v := range c {
// 			if v >= 'A' && v <= 'Z' {
// 				v += 32
// 				temp = append(temp, v)
// 			} else {
// 				temp = append(temp, v)
// 			}
// 		}

// 		for i := 0; i < len(temp); i++ {
// 			if i != len(temp)-1 {
// 				if temp[i+1] == ' ' {
// 					temp[i] -= 32
// 					str += string(temp[i])
// 				} else {
// 					str += string(temp[i])
// 				}
// 			} else {
// 				temp[i] -= 32
// 				str += string(temp[i])
// 			}

// 		}
// 		fmt.Println(str)
// 	}
// }
