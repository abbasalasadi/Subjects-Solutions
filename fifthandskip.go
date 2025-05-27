// allowedFunctions	[ "--allow-builtin", "strings.String" ]

package piscine

func clean(s string) string {
	result := ""
	for _, v := range s {
		if v != ' ' {
			result += string(v)
		}
	}
	return result
}

func FifthAndSkip(str string) string {
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	str = clean(str)
	var list []string
	for len(str) > 5 {
		list = append(list, str[:5])
		str = str[6:]
	}

	if len(str)>0{
		list = append(list, str)
	}

	str = ""
	for i,v:=range list{
		str=str+v
		if i != len(list)-1{
			str = str + " "
		}
	}

	return str+"\n"
}
