// allowedFunctions	[ "--allow-builtin" ]

package piscine

func TrimLast(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			return s[:i]
		}
	}
	return ""
}

func LastWord(s string) string {
	str := TrimLast(s)
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			return s[i:] + "\n"

		}
	}
	return "\n"
}
