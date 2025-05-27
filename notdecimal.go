// allowedFunctions	(3)[ "--allow-builtin", "math.Pow", "strconv.Itoa" ]

package piscine

import (
	"math"
	"strconv"
)

func Valid(s string) bool {
	for _, v := range s {
		if !((v >= '0' && v <= '9') || v == '-' || v == '.') {
			return false
		}
	}
	return true
}

func Neg(s string) (str string, neg bool) {
	if s[0] == '-' {
		neg = true
		str = s[1:]
	} else {
		neg = false
		str = s
	}
	return str, neg
}

func DecimalPoint(str string) (Index int, found bool) {
	for i, v := range str {
		if v == '.' {
			return i, true
		}
	}
	return -1, false
}

func toFlt(num, flt string, neg bool) int {
	x := float64(0)
	y := float64(0)
	result := float64(0)
	for _, c := range num {
		x = x*10 + float64(c-'0')
	}
	for i := len(flt) - 1; i >= 0; i-- {
		y = y/10 + float64(rune(flt[i])-'0')
	}
	result = x + y/10
	if neg {
		result *= -1
	}
		return int(result * math.Pow(10,float64(len(flt))))
}

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	result := 0
	num := ""
	flt := ""
	str, neg := Neg(dec)
	if Valid(dec) {
		Index, found := DecimalPoint(str)
		if found {
			num = str[:Index]
			flt = str[Index+1:]
		} else {
			num = str
		}
		result = toFlt(num, flt, neg)

	} else {
		return dec + "\n"
	}
	return strconv.Itoa(result) + "\n"
}
