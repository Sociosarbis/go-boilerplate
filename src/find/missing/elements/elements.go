package elements

func findMissingElements(nums []int) []int {
	var max int
	min := 101
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	n := max - min + 1
	m := make([]bool, n)
	for _, num := range nums {
		m[num-min] = true
	}
	c := n - len(nums)
	out := make([]int, 0, c)
	for i := min + 1; i < max; i++ {
		if !m[i-min] {
			out = append(out, i)
		}
	}
	return out
}
