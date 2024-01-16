package count

const mask int = 1e9 + 7
const prevIsMaxDigit = 1
const prevLessMaxDigit = 0

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	digits := getDigits(num1)
	count1 := _count(digits, min_sum, max_sum)
	var delta int
	temp := 0
	for i := 0; i < len(digits); i++ {
		temp += digits[i]
	}
	if temp >= min_sum && temp <= max_sum {
		delta = 1
	}
	count2 := _count(getDigits(num2), min_sum, max_sum)
	return (count2 - count1 + delta + mask) % mask
}

func _count(digits []int, min_sum int, max_sum int) int {
	dp := [25][420][2]int{}
	dp[0][0][prevIsMaxDigit] = 1
	n := len(digits)
	for i := 0; i < n; i++ {
		for j := 0; j < 9*n; j++ {
			// 从受限转化为不受限
			for digit := 0; digit < digits[i]; digit++ {
				dp[i+1][j+digit][prevLessMaxDigit] += dp[i][j][prevIsMaxDigit]
				dp[i+1][j+digit][prevLessMaxDigit] %= mask
			}
			// 受限只能从受限转化
			dp[i+1][j+digits[i]][prevIsMaxDigit] += dp[i][j][prevIsMaxDigit]
			dp[i+1][j+digits[i]][prevIsMaxDigit] %= mask
			// 不受限只能转为不受限
			for digit := 0; digit < 10; digit++ {
				dp[i+1][j+digit][prevLessMaxDigit] += dp[i][j][prevLessMaxDigit]
				dp[i+1][j+digit][prevLessMaxDigit] %= mask
			}
		}
	}
	ret := 0
	for i := min_sum; i <= max_sum; i++ {
		ret += dp[n][i][prevIsMaxDigit]
		ret %= mask
		ret += dp[n][i][prevLessMaxDigit]
		ret %= mask
	}
	return ret
}

func getDigits(digits string) []int {
	n := len(digits)
	ret := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ret = append(ret, int(digits[i]-48))
	}
	return ret
}
