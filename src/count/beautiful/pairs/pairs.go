package pairs

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func getFirst(num int) int {
	for num >= 10 {
		num /= 10
	}
	return num % 10
}

func countBeautifulPairs(nums []int) int {
	n := len(nums)
	var ret int
	for i := 0; i < n; i++ {
		a := getFirst(nums[i])
		for j := i + 1; j < n; j++ {
			b := nums[j] % 10
			if gcd(a, b) == 1 {
				ret++
			}
		}
	}
	return ret
}
