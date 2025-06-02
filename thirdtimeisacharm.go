// allowedFunctions	[ "--allow-builtin", "strings.String" ]

package piscine

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 1; i <= len(str)-1; i++ {
		if i%3 == 2 {
			result = result + string(str[i])
		}
	}
	return result + "\n"
}
