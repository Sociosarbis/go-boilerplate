package prime

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func diagonalPrime(nums [][]int) int {
	n := len(nums)
	var ret int
	for i := 0; i < n; i++ {
		if nums[i][i] > ret && isPrime(nums[i][i]) {
			ret = nums[i][i]
		}
		if nums[n-1-i][i] > ret && isPrime(nums[n-1-i][i]) {
			ret = nums[n-1-i][i]
		}
	}
	return ret
}
