package main

import (
	"fmt"
)

func ConcatAlternate(slice1, slice2 []int) []int {
	out := []int{}
	if len(slice1) == len(slice2) {
		for i := range slice1 {
			out = append(out, slice1[i], slice2[i])
		}
	}

	if len(slice1) < len(slice2) {
		end := 0
		for i := range slice1 {
			out = append(out, slice2[i], slice1[i])
			end = i
		}
		out = append(out, slice2[end+1:]...)
	}

	if len(slice1) > len(slice2) {
		end := 0
		for i := range slice2 {
			out = append(out, slice1[i], slice2[i])
			end = i
		}

		out = append(out, slice1[end:]...)
	}

	return out
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
