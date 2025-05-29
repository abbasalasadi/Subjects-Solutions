// allowedFunctions	[ "--allow-builtin" ]

package piscine

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	if value == 0 {
		return "0"
	}

	neg:=false
	if value<0{
		neg = true
		value *= -1
	}
	ref:="0123456789ABCDEF"
	result:=""
	for value > 0 {
		rem:= value%base
		result = string(ref[rem]) + result
		value /= base
	}

	if neg {
		result = "-" + result
	}

	return result
}
