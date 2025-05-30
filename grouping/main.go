// allowedFunctions	(3)[ "fmt.Printf", "os.Args", "--allow-builtin" ]

package main

import (
	"fmt"
	"os"
)

func ValidExp(e string) ([]string, bool) {
	if e[0] == '(' && e[len(e)-1] == ')' {
		exp := e[1 : len(e)-1]
		for _, c := range exp {
			if !((c >= 'a' && c <= 'z') || c == '|') {
				return nil, false
			}
		}
		if exp[0] == '|' || exp[len(exp)-1] == '|' {
			return nil, false
		}
		exps := []string{}
	outer:
		for len(exp) > 0 {
			for i, v := range exp {
				if v == '|' {
					exps = append(exps, exp[:i])
					exp = exp[i+1:]
					break
				}
				if i == len(exp)-1 {
					exps = append(exps, exp)
					break outer
				}
			}
		}
		return exps, true
	}
	return nil, false

}

func contains(str, substr string) bool {
	if len(substr) == 0 {
		return true 
	}

	for i := 0; i <= len(str)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if str[i+j] != substr[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func split(s string) (words []string) {
outer:
	for len(s) > 0 {
		for i, c := range s {
			if c == ' ' || c == '.' || c == ',' {
				words = append(words, string(s[:i]))
				s = s[i+1:]
				break
			}
			if i == len(s)-1 {
				words = append(words, string(s))
				break outer
			}
		}
	}
	return words
}

func And(a []bool) bool {
	for _, c := range a {
		if c == false {
			return false
		}
	}
	return true
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 || arg[1] == "" {
		return
	}
	words := split(arg[1])
	exps, valid := ValidExp(arg[0])
	if valid {
		i := 0
		for _, word := range words {
			for _, exp := range exps {
				if contains(word, exp) {
					i++
					fmt.Printf("%v: %v\n", i, word)
				}
			}
		}
	} else {
		return
	}
}
