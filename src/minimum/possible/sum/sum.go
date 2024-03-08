package sum

const mask int64 = 1e9 + 7

func minimumPossibleSum(n int, target int) int {
	n64 := int64(n)
	target64 := int64(target)
	half := target64 / 2
	if half > n64 {
		half = n64
	}
	ret := ((1 + half) * half / 2)
	remain := n64 - half
	if remain > 0 {
		ret += (target64*2 + remain - 1) * remain / 2
	}
	return int(ret % mask)
}
