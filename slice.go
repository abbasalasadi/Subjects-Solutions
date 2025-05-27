package piscine

func Rev(n,max int) (num int) {
	if n >= 0 {
			num = n
		} else {
			num = max + n+1
		}
		return num
	}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) > 2 {
		return nil
	}
	var start int
	var end int
	if len(nbrs) == 1{
		start = Rev(nbrs[0],len(a)-1)
		end = len(a)-1
	}
	if len(nbrs) == 2 {
		start = Rev(nbrs[0],len(a)-1)
		end = Rev(nbrs[1],len(a)-1)-1
		if start > end {
			return nil
		}
	}
	return a[start:end+1]
}
