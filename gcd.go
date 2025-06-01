package piscine

func Gcd(a, b uint) uint {
	var n int
	if a > b {
		n = int(b)
	} else {
		n = int(a)
	}
	result := 0
	for i := 1; i <= n; i++ {
		if int(a)%i == 0 && int(b)%i == 0 {
			result = i
		}
	}
	return uint(result)
}
