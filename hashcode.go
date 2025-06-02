// allowedFunctions	[ "--allow-builtin" ]

package piscine

func HashCode(dec string) string {
	str := ""
	for _, c := range dec {
		str = str + string((c + rune(len(dec))%127))
	}
	return str
}
