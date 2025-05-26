// allowedFunctions	[ "--allow-builtin" ]

package piscine

func CanJump(j []uint) bool {
	var Pos uint = 0
	var Jump uint = j[0]
	if Pos == 0 && Jump == 0 {
		return true
	}
	for {
		if Jump == 0 {
			return false
		}
		Pos = Pos + Jump
		if Pos == uint(len(j)-1){
			return true
		}
		Jump = j[Pos]
	}
}
