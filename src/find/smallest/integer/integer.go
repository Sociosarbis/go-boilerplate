package integer

func findSmallestInteger(nums []int, value int) int {
	counter := make([]int, value)
	for _, num := range nums {
		if num < 0 {
			num += (((-num) / value) + 1) * value
		}
		counter[num%value]++
	}
	var i int
	for {
		index := i % value
		if counter[index] > 0 {
			i++
			counter[index]--
		} else {
			break
		}
	}
	return i
}
