// allowedFunctions	[ "len", "--allow-builtin" ]

package piscine

func Itoa(n int) string {
	neg := false
	if n < 0 {
		neg = true
		n *= -1
	}

	str := ""
	if n == 0 {
		return "0"
	}
	for n > 0 {
		str = string(n%10+'0') + str
		n /= 10
	}
	if neg {
		str = "-" + str
	}
	return str
}
