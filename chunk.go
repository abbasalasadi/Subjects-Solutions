// allowedFunctions	[ "fmt.Println", "--allow-builtin" ]

package piscine

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 || len(slice) == 0 {
		fmt.Println()
		return
	}

	col := len(slice) / size
	if len(slice)%size > 0 {
		col += 1
	}
	list := [][]int{}
	fmt.Println(len(list))

	for len(slice) > 0 {
		if len(slice) < size {
			size = len(slice)
		}

		list = append(list, slice[:size])
		slice = slice[size:]
	}
	fmt.Println(list)
}
