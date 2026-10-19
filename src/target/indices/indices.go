package indices

func targetIndices(nums []int, target int) []int {
	left := -1
	var count int
	for _, num := range nums {
		if num < target {
			left++
		} else if num == target {
			count++
		}
	}
	if count == 0 {
		return nil
	}
	left++
	out := make([]int, 0, count)
	for i := 0; i < count; i++ {
		out = append(out, left+i)
	}
	return out
}
