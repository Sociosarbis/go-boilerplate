package disjoint

func partitionDisjoint(nums []int) int {
	arr := make([]int, 1e6+1)
	var min int = 1e6
	for _, num := range nums {
		arr[num]++
		if num < min {
			min = num
		}
	}

	max := 0
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		arr[nums[i]]--
		if nums[i] == min && arr[min] == 0 {
			for j := min + 1; j < len(arr); j++ {
				if arr[j] != 0 {
					min = j
					break
				}
			}
		}
		if max <= min {
			return i + 1
		}
	}
	return 0
}
