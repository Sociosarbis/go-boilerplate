package function

func maxRotateFunction(nums []int) int {
	max := 0
	sum := 0
	for i, num := range nums {
		max += i * num
		sum += num
	}

	prev := max
	n := len(nums)

	for i := 1; i < n; i++ {
		temp := prev - n*nums[n-i] + sum
		if temp > max {
			max = temp
		}
		prev = temp
	}
	return max
}
