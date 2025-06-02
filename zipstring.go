// allowedFunctions	[ "--allow-builtin", "strconv.Itoa" ]

package piscine

import (
	"strconv"
)

func ZipString(s string) string {
	count := 1
	str := ""
	n := ""
	for i := 0; i <= len(s)-2; i++ {
		if i != len(s)-2 {
			if s[i] == s[i+1] {
				count++
				continue
			}
			n = strconv.Itoa(count)
			str = str + n + string(s[i])
			count = 1
		} else {
			if s[i] == s[i+1] {
				count++
				n = strconv.Itoa(count)
				str = str + n + string(s[i])
			} else {
				n = strconv.Itoa(count)
				str = str + n + string(s[i]) + "1" + string(s[i+1])
			}
		}
	}
	return str
}
