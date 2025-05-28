// allowedFunctions	(5)[ "--allow-builtin", "fmt.*", "os.Args", "strconv.Atoi", "strings.*" ]

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arr(s string) ([]int, error) {
	if s[0] != '[' || s[len(s)-1] != ']' {
		return []int{}, fmt.Errorf("Invalid input.")
	}
	str := strings.Trim(s, "[]")
	StrList := strings.Split(str, ",")
	NumList := []int{}
	for _, v := range StrList {
		st := strings.TrimSpace(v)
		num, err := strconv.Atoi(st)
		if err != nil {
			return []int{}, fmt.Errorf("Invalid number: p")
		}
		NumList = append(NumList, num)
	}
	return NumList, nil
}

func FindPairs(Nums []int, targetSum int) (List [][]int) {
outer:
	for i:=0; len(Nums) > 0; i++ {
		a:=Nums[0]
		Nums = Nums [1:]
		for j,v:=range Nums{
			if targetSum == a + v {
				List=append(List,[]int{i,i+j+1})
				Nums=append(Nums[:j], Nums[j:]...)
				break
			}
			// fmt.Println(i,j,Nums[0] , v, Nums)
	}
	if i==len(Nums){
		break outer
	}
	}
	return List
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("Invalid input.")
		return
	}

	targetSum, err := strconv.Atoi(arg[1])
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}
	NumList, err := arr(arg[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(NumList)
	// fmt.Println("targetSum: ", targetSum)
	List:=FindPairs(NumList,targetSum)
	if len(List)==0{
		fmt.Println("No pairs found.")
	} else {
	fmt.Printf("Pairs with sum %v: %v",targetSum, FindPairs(NumList,targetSum))
	}
}
