// allowedFunctions	[ "--allow-builtin", "strings.Contains" ]

package piscine

func FindUnique(str1, str2 string) string {
	found := false
	unique := ""
	for i, c := range str1 {
		for j, v := range str2 {
			if i != j {
				if c == v {
					found = true
					break
				}
			}
		}
		if !found {
			unique = unique + string(c)
		}
		found = false
	}
	return unique
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	return len(FindUnique(str1, str2)) + len(FindUnique(str2, str1))
}
