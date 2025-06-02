// allowedFunctions	[ "--allow-builtin" ]

package piscine

func IsAlpha(r rune) bool {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return true
	}
	return false
}

func IsCapital(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
func IsSmall(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsCamel(s string) bool {
	for i := range s {
		if !IsAlpha(rune(s[i])) {
			return false
		}
	}
	for i := 0; i <= len(s)-2; i++ {
		if IsCapital(rune(s[i])) && IsCapital(rune(s[i+1])) {
			return false
		}
	}
	return true
}

func CamelToSnakeCase(s string) string {
	if IsCamel(s) {
		str := ""
	outer:
		for len(s) > 0 {
			for i := 0; i <= len(s)-2; i++ {
				if IsSmall(rune(s[i])) && IsCapital(rune(s[i+1])) {
					str = str + s[:i+1] + "_"
					s = s[i+1:]
					break
				}
				if i == len(s)-2 {
					str = str + s
					break outer
				}
			}
		}
		return str
	}
	return s
}
