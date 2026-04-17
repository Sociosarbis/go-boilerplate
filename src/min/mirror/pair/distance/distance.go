package distance

func reverse(num int) int {
	var out int
	for num != 0 {
		d := num % 10
		out = out*10 + d
		num /= 10
	}
	return out
}

func normalize(num int) int {
	for num%10 == 0 {
		num /= 10
	}
	return num
}

func minMirrorPairDistance(nums []int) int {
	n := len(nums)
	numToIndex := make(map[int]int, n)
	out := n
	for i, num := range nums {
		if j, ok := numToIndex[num]; ok {
			if i-j < out {
				out = i - j
			}
		}
		num = normalize(num)
		numToIndex[reverse(num)] = i
	}
	if out == n {
		return -1
	}
	return out
}
