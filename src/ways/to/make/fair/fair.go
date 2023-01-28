package fair

func waysToMakeFair(nums []int) int {
	p1 := make([]int, len(nums))
	p2 := make([]int, len(nums))

	for i, num := range nums {
		if i&1 == 0 {
			if i == 0 {
				p1[i] = num
			} else {
				p2[i] = p2[i-1]
				p1[i] = p1[i-2] + num
			}
		} else {
			p1[i] = p1[i-1]
			if i == 1 {
				p2[i] = num
			} else {
				p2[i] = p2[i-2] + num
			}
		}
	}
	ret := 0
	for i := range nums {
		var left1 int
		var left2 int
		if i > 0 {
			left1 = p1[i-1]
			left2 = p2[i-1]
		}

		var right1 int
		var right2 int
		if i+1 < len(nums) {
			right1 = p2[len(nums)-1] - p2[i]
			right2 = p1[len(nums)-1] - p1[i]
		}
		if left1+right1 == left2+right2 {
			ret++
		}
	}

	return ret
}
