// allowedFunctions	[ "--allow-builtin" ]

package piscine

func TrimStart(s string) string {
	for i,c:=range s{
		if c != ' '{
			return s[i:]
		}
	}
	return ""
}

func FirstWord(s string) string {
	str:=TrimStart(s)
	for i, c := range str {
		if c == ' ' {
			return str[:i] + "\n"
		}
	}
	return "\n"
}
