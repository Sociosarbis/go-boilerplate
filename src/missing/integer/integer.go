package integer

func missingInteger(nums []int) int {
	var sum, count int
	tempSum, temp := nums[0], 1
	visited := [51]bool{}
	for _, num := range nums {
		visited[num] = true
	}
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			temp++
			tempSum += nums[i]
		} else {
			break
		}
	}
	if temp > count {
		sum = tempSum
	}
	for sum <= 50 {
		if !visited[sum] {
			break
		}
		sum++
	}
	return sum
}
