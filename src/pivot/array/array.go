package array

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	out := make([]int, n)
	var j int
	for _, num := range nums {
		if num < pivot {
			out[j] = num
			j++
		}
	}
	for _, num := range nums {
		if num == pivot {
			out[j] = num
			j++
		}
	}
	for _, num := range nums {
		if num > pivot {
			out[j] = num
			j++
		}
	}
	return out
}
