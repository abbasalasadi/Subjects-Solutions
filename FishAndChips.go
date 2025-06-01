package piscine

func div2(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func div3(n int) bool {
	if n%3 == 0 {
		return true
	}
	return false
}

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	div2 := div2(n)
	div3 := div3(n)

	if div2 && div3 {
		return "fish and chips"
	} else if div2 {
		return "fish"
	} else if div3 {
		return "chips"
	}
	return "error: non divisible"
}
