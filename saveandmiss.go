// allowedFunctions	[ "--allow-builtin" ]

package piscine

func SaveAndMiss(arg string, num int) string {
	result := ""
	Save := true
	count := 0
	for _, v := range arg {
		if Save {
			result += string(v)
		}
		count++
		if count == num {
			count = 0
			Save = !Save
		}
	}
	return result
}
