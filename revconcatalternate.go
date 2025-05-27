// allowedFunctions	[ "--allow-builtin" ]

package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	var rem []int
	switch {
	case len(slice1) > len(slice2):
		rem = slice1[len(slice2):]
		slice1 = slice1[:len(slice2)]
	case len(slice1) < len(slice2):
		rem = slice2[len(slice1):]
		slice2 = slice2[:len(slice1)]
	}

	var result []int

		for i := len(rem) - 1; i >= 0; i-- {
			result = append(result, rem[i])
		}

	for i := len(slice1) - 1; i >= 0; i-- {
		result = append(result, slice1[i], slice2[i])
	}
	return result
}
