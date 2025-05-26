// allowedFunctions	(4)[ "strconv.Atoi", "os.Args", "fmt.*", "--allow-builtin" ]

package main

import (
	"fmt"
	"os"
	"strconv"
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
		if n%i == 0 || n%(n+2) == 0 {
			return false
		}
	}
	return true
}

func NextPrime(n int) (int) {
	i:=n+1
	for {
		if isPrime(i) {
			return i
		}
		i++
	}
}

func main() {

	arg:=os.Args[1:]
	if len(arg)!=1{return}
	num,err:=strconv.Atoi(arg[0])
	if err!=nil{
		return
	}
	Prime:=2
	PrimeList:=[]int{}
	for num>0{
		if num%Prime == 0{
			PrimeList = append(PrimeList,Prime)
			num /= Prime
			continue
		}
		Prime = NextPrime(Prime)
		if Prime > num {
			break
		}
	}
	for i,v:=range PrimeList{
	fmt.Print(v)
	if i != len(PrimeList)-1{
		fmt.Print("*")
	}
	}
}
