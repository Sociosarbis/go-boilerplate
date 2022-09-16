package area

func helper(all [][]int, cur []int, start int) [][]int {
	if start >= len(all) {
		return append(all, cur)
	}
	rec := all[start]
	if cur[2] <= rec[0] || cur[3] <= rec[1] || cur[0] >= rec[2] || cur[1] >= rec[3] {
		return helper(all, cur, start+1)
	}
	if cur[0] < rec[0] {
		all = helper(all, []int{cur[0], cur[1], rec[0], cur[3]}, start+1)
	}
	if cur[2] > rec[2] {
		all = helper(all, []int{rec[2], cur[1], cur[2], cur[3]}, start+1)
	}
	if cur[1] < rec[1] {
		var left int
		if rec[0] > cur[0] {
			left = rec[0]
		} else {
			left = cur[0]
		}
		var right int
		if rec[2] > cur[2] {
			right = rec[2]
		} else {
			right = cur[2]
		}
		all = helper(all, []int{left, cur[1], right, rec[1]}, start+1)
	}
	if cur[3] > rec[3] {
		var left int
		if rec[0] > cur[0] {
			left = rec[0]
		} else {
			left = cur[0]
		}
		var right int
		if rec[2] > cur[2] {
			right = rec[2]
		} else {
			right = cur[2]
		}
		all = helper(all, []int{left, rec[3], right, cur[3]}, start+1)
	}
	return all
}

func rectangleArea(rectangles [][]int) int {
	var res int64
	var M int64 = 1e9 + 7
	all := [][]int{}
	for _, rectangle := range rectangles {
		all = helper(all, rectangle, 0)
	}
	for _, a := range all {
		res = (res + int64(a[2]-a[0])*int64(a[3]-a[1])) % M
	}
	return int(res)
}
