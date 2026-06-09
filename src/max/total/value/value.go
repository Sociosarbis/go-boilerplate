package value

const mx int = 1e9

func maxTotalValue(nums []int, k int) int64 {
	var max int
	min := mx
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return int64(max-min) * int64(k)
}
