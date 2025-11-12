package operations13

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func minOperations(nums []int) int {
	var ones int
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}
	n := len(nums)
	if ones != 0 {
		return n - ones
	}
	i := 2
	for ; i <= n; i++ {
		end := n - i + 1
		for j := 0; j < end; j++ {
			temp := nums[j]
			for k := 1; k < i; k++ {
				temp = gcd(temp, nums[j+k])
			}
			if temp == 1 {
				return n + i - 2
			}
		}
	}
	return -1
}
