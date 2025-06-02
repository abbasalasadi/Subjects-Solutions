// allowedFunctions	[ "--allow-builtin" ]

package piscine

func Split(s string) (arr []string) {
outer:
	for {
		for i, c := range s {
			if i == len(s)-1 {
				arr = append(arr, s)
				break outer
			}
			if c == ' ' {
				arr = append(arr, s[:i])
				s = s[i+1:]
				break
			}
		}
	}
	return arr
}

func Check(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return false
	}
	return true
}

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	arr := Split(s)
	for _, c := range arr {
		if !Check(rune(c[0])) {
			return false
		}
	}
	return true
}
