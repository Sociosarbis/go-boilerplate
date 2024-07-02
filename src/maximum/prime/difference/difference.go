package difference

func isPrime(num int) bool {
	if num < 8 {
		return num == 2 || num == 3 || num == 5 || num == 7
	}
	return num%2 != 0 && num%3 != 0 && num%5 != 0 && num%7 != 0
}

func maximumPrimeDifference(nums []int) int {
	n := len(nums)
	var i int
	for ; i < n; i++ {
		if isPrime(nums[i]) {
			break
		}
	}
	j := n - 1
	for ; j > i; j-- {
		if isPrime(nums[j]) {
			break
		}
	}
	return j - i
}
