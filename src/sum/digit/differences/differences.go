package differences

func sumDigitDifferences(nums []int) int64 {
	var n int
	temp := nums[0]
	for temp != 0 {
		n++
		temp /= 10
	}
	counter := make([][10]int, n)
	for _, num := range nums {
		var i int
		for num != 0 {
			counter[i][num%10]++
			num /= 10
			i++
		}
	}

	var ret int64
	for _, row := range counter {
		for i := 0; i < 10; i++ {
			if row[i] != 0 {
				for j := i + 1; j < 10; j++ {
					if row[j] != 0 {
						ret += int64(row[i]) * int64(row[j])
					}
				}
			}
		}
	}
	return ret
}
