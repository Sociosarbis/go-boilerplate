package coprimes

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func replaceNonCoprimes(nums []int) []int {
	var i int
	n := len(nums)
	for j := 1; j < n; j++ {
		g := gcd(nums[i], nums[j])
		if g == 1 {
			nums[i+1] = nums[j]
			i++
		} else {
			nums[i] *= nums[j] / g
			for i > 0 {
				g = gcd(nums[i], nums[i-1])
				if g == 1 {
					break
				}
				nums[i-1] *= nums[i] / g
				i--
			}
		}
	}
	nums = nums[:i+1]
	return nums
}
