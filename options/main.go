// allowedFunctions	(3)[ "os.*", "fmt.*", "--allow-builtin" ]
package main

import (
	"fmt"
	"os"
)

func Trim(s string) (str string) {
	for i, v := range s {
		if v != ' ' && i != len(str)-1 {
			str = s[i:]
			break
		} else {
			str = s
		}
	}
	for i, v := range str {
		if v == ' ' || i == len(str)-1 {
			if v == ' ' {
				str = str[:i]
			} else {
				str = str[:i+1]
			}
			break
		}
	}
	return str
}

func valid(s string) bool {
	if s == "-"{
		return false
	}
	for _,v:=range s {
		if !((v>='a' && v<='z') || v=='-') {
			return false
		}
	}
	return true
}

func all(arg []string) (string, bool) {
	if len(arg) == 0 {
		return "", true
	}
	for _, v := range arg {
		if v[0] == '-' && v[1] == 'h' {
			return "", true
		}
	}
	str := ""
	for _, v := range arg {
		str = str + v
	}
	return str, false
}

func main() {
	arg := os.Args[1:]
	for i, v := range arg {
		arg[i] = Trim(v)
		if !valid(arg[i]){
			fmt.Println("Invalid Option")
			return
		}
	}
	str, all := all(arg)
	if all {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	m := make(map[rune]int)
	for i := '['; i <= 'z'; i++ {
		m[i] = 0
	}
	for _,v:=range str{
		m[v]=1
	}
	options:="012345zyxwvutsrqponmlkjihgfedcba"
	for i,v:=range options{
		fmt.Print(m[v])
		if (i+1)%8==0{
			fmt.Print(" ")
		}

	}

}
