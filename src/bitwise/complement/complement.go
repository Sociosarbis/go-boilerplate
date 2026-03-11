package complement

func bitwiseComplement(n int) int {
	mask := 1
	temp := n
	for temp > 1 {
		temp >>= 1
		mask <<= 1
	}
	return mask - 1 + mask - n
}
