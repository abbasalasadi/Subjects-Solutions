// allowedFunctions	[ "--allow-builtin", "strings.String" ]

package piscine

func RetainFirstHalf(str string) string {
	if str == "" {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	center := len(str) / 2
	return str[:center]
}
