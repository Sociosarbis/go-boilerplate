package score

func maxDivScore(nums []int, divisors []int) int {
	var ret int
	count := -1
	var temp int
	for _, d := range divisors {
		for _, num := range nums {
			if num%d == 0 {
				temp++
			}
		}
		if temp > count {
			ret = d
			count = temp
		} else if temp == count && d < ret {
			ret = d
		}
		temp = 0
	}
	return ret
}
