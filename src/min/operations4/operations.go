package operations4

func minOperations(nums []int, x int) int {
	m := map[int]int{}
	temp := 0
	ret := -1
	for i := len(nums) - 1; i >= 0; i-- {
		temp += nums[i]
		if temp <= x {
			m[temp] = len(nums) - i
			if temp == x {
				ret = len(nums) - i
			}
		} else {
			break
		}
	}
	temp = 0
	for i, num := range nums {
		temp += num
		if temp > x {
			break
		}
		if temp == x {
			if i+1 < ret || ret == -1 {
				ret = i + 1
			}
		} else if c, ok := m[x-temp]; ok {
			if i+1+c <= len(nums) && (ret == -1 || i+1+c < ret) {
				ret = i + 1 + c
			}
		}
	}
	return ret
}
