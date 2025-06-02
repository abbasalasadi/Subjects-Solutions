// allowedFunctions	[ "--allow-builtin" ]

package piscine

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(n+2) == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for i := nb; i >= 0; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}
