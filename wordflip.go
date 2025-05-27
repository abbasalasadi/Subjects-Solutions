// allowedFunctions	(3)[ "--allow-builtin", "strings.Split", "strings.TrimSpace" ]

package piscine

import (
	"strings"
)

func WordFlip(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	}

	list := strings.Split(str, " ")
	var clean []string
	for i:=len(list)-1; i>=0; i-- {
		list[i] = strings.TrimSpace(list[i])
		if list[i] != "" {
			clean = append(clean, list[i])
		}
	}

	var result string
	for i,v:=range clean{
		result = result + v
		if i != len(clean)-1 {
			result = result + " "
		}
	}
	return result + "\n"
}
