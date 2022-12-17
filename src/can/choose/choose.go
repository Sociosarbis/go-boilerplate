package choose

func canChoose(groups [][]int, nums []int) bool {
	j := 0
loop:
	for i := 0; i < len(groups); i++ {
		for ; j < len(nums); j++ {
			if nums[j] == groups[i][0] {
				k := 1
				for ; k < len(groups[i]); k++ {
					if nums[j+k] != groups[i][k] {
						break
					}
				}
				if k == len(groups[i]) {
					j += k
					continue loop
				}
			}
		}
		return false
	}
	return true
}
