package operations7

func minOperations(nums []int) int {
	var hasOne bool
	var hasZero bool
	n := len(nums)
	var ret int
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 0 {
			if !hasZero {
				hasZero = true
			}
		} else {
			if hasZero {
				if hasOne {
					ret += 2
				} else {
					ret++
				}
				hasZero = false
			}
			if !hasOne {
				hasOne = true
			}
		}
	}
	if hasZero {
		if hasOne {
			ret += 2
		} else {
			ret++
		}
	}
	return ret
}
