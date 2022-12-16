package elements

func minElements(nums []int, limit int, goal int) int {
	goal64 := int64(goal)
	for _, num := range nums {
		goal64 -= int64(num)
	}
	if goal64 < 0 {
		goal64 = -goal64
	}
	limit64 := int64(limit)
	if goal64%limit64 == 0 {
		return int(goal64 / limit64)
	} else {
		return int(goal64/limit64) + 1
	}
}
