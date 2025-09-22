package elements

func maxFrequencyElements(nums []int) int {
	counter := [101]int{}
	var max int
	for _, num := range nums {
		counter[num]++
		if counter[num] > max {
			max = counter[num]
		}
	}
	var out int
	for i := 1; i <= 100; i++ {
		if counter[i] == max {
			out += max
		}
	}
	return out
}
