package group

func countLargestGroup(n int) int {
	var counter [29]int
	for i := 1; i <= n; i++ {
		var sum int
		v := i
		for v != 0 {
			sum += v % 10
			v /= 10
		}
		counter[sum]++
	}
	var ret int
	var max int
	for i := 1; i < 29; i++ {
		if counter[i] > max {
			ret = 1
			max = counter[i]
		} else if counter[i] == max {
			ret++
		}
	}
	return ret
}
