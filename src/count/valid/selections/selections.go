package selections

func countValidSelections(nums []int) int {
	var count int
	for _, num := range nums {
		if num != 0 {
			count += num
		}
	}
	var temp, out int
	for _, num := range nums {
		if num == 0 {
			back := count - temp
			if back == temp {
				out += 2
			} else if back == temp+1 || back == temp-1 {
				out++
			}
		} else {
			temp += num
		}
	}
	return out
}
