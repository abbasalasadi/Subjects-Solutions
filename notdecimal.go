// allowedFunctions	(3)[ "--allow-builtin", "math.Pow", "strconv.Itoa" ]

package piscine

import (
	// "fmt"

	"math"
	"strconv"
)

func Float(str string) (Index int, Err bool) {
	Err = false
	Index = -1
	for i, v := range str {
		if v == '.' {
			Index = i
		}
		if !((v >= '0' && v <= '9') || v == '.') {
			Err = true
			Index = 0
			break
		}
	}
	return Index, Err
}

func ConvFloat(str string) (num float64, Index int, Err bool) {
	Index, Err = Float(str)
	if Err {
		return 0, 0, true
	}
	num = float64(0)
	Int := ""
	dec := ""
	if Index != -1 {
		Int = str[:Index]
		dec = str[Index+1:]
		for _, v := range Int {
			num = num*10 + float64(v-'0')
		}
		dot := float64(0)
		for i := len(dec) - 1; i >= 0; i-- {
			dot = dot/10 + float64(dec[i]-'0')
		}
		return num + dot/10, Index, false
	}
	Int = str
	for _, v := range Int {
		num = num*10 + float64(v-'0')
	}
	return num, Index, false
}

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}
	str := ""
	Neg := false
	if dec[0] == '-' {
		str = dec[1:]
		Neg = true
	} else {
		str = dec
	}
	num, Index, Err := ConvFloat(str)
	if Err {
		if !Neg {
			return str + "\n"
		}
		return "-" + str + "\n"
	}
	result := 0.0
	if Index != -1 {
		result = num * math.Pow(10, float64(len(str)-Index-1))
	} else {
		result = num
	}
	if Neg {
		result *= -1
	}
	return strconv.Itoa(int(result)) + "\n"
}
