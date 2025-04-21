package arrays

const max int = 3 * 1e5
const min int = -3 * 1e5

func numberOfArrays(differences []int, lower int, upper int) int {
	var l int
	var r int
	var temp int
	for _, num := range differences {
		temp += num
		if temp < min {
			return 0
		}
		if temp > max {
			return 0
		}
		if temp < l {
			l = temp
		}
		if temp > r {
			r = temp
		}
	}
	l = lower - l
	r = upper - r
	if l <= r {
		return r - l + 1
	}
	return 0
}
