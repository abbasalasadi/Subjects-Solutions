// "github.com/01-edu/z01.PrintRune", "os.*", "--allow-builtin"

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	num := 0
	for _, v := range s {
		num = num*10 + int(v-'0')
	}
	return num
}

func Print(n int) {
	for n > 0 {
		rem := n % 10
		defer z01.PrintRune(rune(rem) + '0')
		n /= 10

	}
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 || Atoi(arg[0]) < 0 {
		z01.PrintRune('0')
		return
	}
	num := Atoi(arg[0])
	Sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			Sum += i
		}
	}

	Print(Sum)

}
