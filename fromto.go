package piscine

import (
	"strconv"
)

func FromTo(from int, to int) string {
	num := ""
	out := ""
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid"
	}
	if from < to {
		for i := from; i <= to; i++ {
			num = strconv.Itoa(i)
			if i < 10 {
				num = "0" + num
			}
			out = out + num
			if i != to {
				out = out + ", "
			}
		}
	}
	if from > to {
		for i := from; i >= to; i-- {
			num = strconv.Itoa(i)
			if i < 10 {
				num = "0" + num
			}
			out = out + num
			if i != to {
				out = out + ", "
			}
		}
	}
	if from == to {
		num = strconv.Itoa(from)
		if from < 10 {
			num = "0" + num
		}
		out = out + num
	}
	return out
}
