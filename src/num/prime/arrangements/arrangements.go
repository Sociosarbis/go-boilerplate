package arrangements

var primNumsLessThan100 = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func numPrimeArrangements(n int) int {
	count := 0
	m := 1000000007
	for i := range primNumsLessThan100 {
		if primNumsLessThan100[i] > n {
			count = i
			break
		}
	}
	ret := 1
	for i := 2; i <= count; i++ {
		ret *= i
		if ret >= m {
			ret %= m
		}
	}
	count = n - count
	for i := 2; i <= count; i++ {
		ret *= i
		if ret >= m {
			ret %= m
		}
	}
	return ret
}
