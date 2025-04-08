package operations2

func minimumOperations(nums []int) int {
	counter := [101]int{}
	var duplicateNum int
	for _, num := range nums {
		if counter[num] == 1 {
			duplicateNum++
		}
		counter[num]++
	}
	n := len(nums)
	for i := 0; i < n; i += 3 {
		for j := 0; j < 3 && i+j < n; j++ {
			counter[nums[i+j]]--
			if counter[nums[i+j]] == 1 {
				duplicateNum--
				if duplicateNum == 0 {
					return (i / 3) + 1
				}
			}
		}
	}
	return 0
}
