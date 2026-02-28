package binary

func concatenatedBinary(n int) int {
	digits := [17]int{}
	var out int
	for i := 1; i <= n; i++ {
		var count int
		temp := i
		for temp > 0 {
			digits[count] = temp & 1
			temp >>= 1
			count++
		}
		// 秦九韶算法
		// (a + b) % m = (a % m + b % m) % m
		// (a * b) % m = ((a % m) * (b % m)) % m
		for j := count - 1; j >= 0; j-- {
			out = (out*2 + digits[j]) % 1000000007
		}
	}
	return out
}
