// allowedFunctions	[ "--allow-builtin" ]

package piscine

func RepeatAlpha(s string) string {
	itr := []int{}
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			itr = append(itr, int(c-'a'+1))
		} else if c >= 'A' && c <= 'Z' {
			itr = append(itr, int(c-'A'+1))
		} else {
			itr = append(itr, 1)
		}
	}
	str := ""
	for i, c := range s {
		for j := 0; j < itr[i]; j++ {
			str = str + string(c)
		}
	}
	return str
}
