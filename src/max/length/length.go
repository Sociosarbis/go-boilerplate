package length

func maxLength(arr []string) int {
	masks := make([]int, len(arr))
	for i, str := range arr {
		mask := 0
		for _, char := range str {
			bit := 1 << (char - 97)
			if bit&mask == 0 {
				mask |= bit
			} else {
				mask = 0
				break
			}
		}
		masks[i] = mask
	}
	dp := map[int]int{
		0: 0,
	}
	ret := 0
	for i, mask := range masks {
		if mask != 0 {
			newDp := map[int]int{}
			for bit, length := range dp {
				newDp[bit] = length
				if bit&mask == 0 {
					newBit := bit | mask
					newLength := length + len(arr[i])
					newDp[newBit] = newLength
					if newLength > ret {
						ret = newLength
					}
				}
			}
			dp = newDp
		}
	}
	return ret
}
