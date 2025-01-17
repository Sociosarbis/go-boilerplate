package length2

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	n := len(nums)
	ret := n + 1
	var temp int
	counter := make([]int, 31)
	var j int
	for i := 0; i < n; i++ {
		if j < i {
			j = i
		}
		for temp < k && j < n {
			num := nums[j]
			temp |= num
			for i := 0; i < 31; i++ {
				if num&(1<<i) != 0 {
					counter[i]++
				}
			}
			j++
		}
		if temp >= k {
			if j-i < ret {
				ret = j - i
			}
			num := nums[i]
			temp |= num
			for j := 0; j < 31; j++ {
				if num&(1<<j) != 0 {
					counter[j]--
					if counter[j] == 0 {
						temp -= 1 << j
					}
				}
			}
		} else {
			break
		}
	}
	if ret == n+1 {
		return -1
	}
	return ret
}
