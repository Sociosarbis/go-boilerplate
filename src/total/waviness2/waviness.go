package waviness2

const s int = 4

func revert(digits []int) []int {
	var l int
	r := len(digits) - 1
	for l < r {
		digits[l], digits[r] = digits[r], digits[l]
		l++
		r--
	}
	return digits
}

func getDigits(num int64) []int {
	out := make([]int, 0, 16)
	for num != 0 {
		out = append(out, int(num%10))
		num /= 10
	}
	return revert(out)
}

func isWave(l, k, m int) bool {
	return (k > l && k > m) || (k < l && k < m)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type item struct {
	count int
	value int64
}

func subTotalWaviness(digits1 []int, digits2 []int) int64 {
	n := len(digits1)
	// position, tight, last, secondLast
	dp := make([][][][]item, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][][]item, s)
		for j := 0; j < s; j++ {
			dp[i][j] = make([][]item, 10)
			for k := 0; k < 10; k++ {
				dp[i][j][k] = make([]item, 10)
			}
		}
	}
	for i := 0; i < s; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				dp[0][i][j][k].value = -1
			}
		}
	}
	for i := digits1[0]; i <= digits2[0]; i++ {
		var j int
		if i == digits1[0] {
			j |= 1
		}
		if i == digits2[0] {
			j |= 2
		}
		dp[0][j][i][i].count, dp[0][j][i][i].value = 1, 0
	}
	for i := 1; i < n; i++ {
		index := i & 1
		prevIndex := 1 - index
		for j := 0; j < s; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					dp[index][j][k][l].count, dp[index][j][k][l].value = 0, -1
				}
			}
		}
		for j := 0; j < s; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					if dp[prevIndex][j][k][l].value != -1 {
						for m := 0; m < 10; m++ {
							if (j&2 != 0 && m > digits2[i]) || (j&1 != 0 && m < digits1[i]) {
								continue
							}
							var nextJ int
							if m == digits1[i] && j&1 != 0 {
								nextJ |= 1
							}
							if m == digits2[i] && j&2 != 0 {
								nextJ |= 2
							}
							if dp[index][nextJ][m][k].value == -1 {
								dp[index][nextJ][m][k].value = 0
							}
							dp[index][nextJ][m][k].value += dp[prevIndex][j][k][l].value
							if isWave(l, k, m) {
								dp[index][nextJ][m][k].value += int64(dp[prevIndex][j][k][l].count)
							}
							dp[index][nextJ][m][k].count += dp[prevIndex][j][k][l].count
						}
					}
				}
			}
		}
	}
	var out int64
	index := (n - 1) & 1
	for i := 0; i < s; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if dp[index][i][j][k].value > 0 {
					out += dp[index][i][j][k].value
				}
			}
		}
	}
	return out
}

func totalWaviness(num1 int64, num2 int64) int64 {
	digits1, digits2 := getDigits(num1), getDigits(num2)
	n1, n2 := len(digits1), len(digits2)
	var out int64
	if n1 < n2 {
		target := make([]int, n1)
		for i := 0; i < n1; i++ {
			target[i] = 9
		}
		out += subTotalWaviness(digits1, target)
	}
	for n1 < n2 {
		digits1 = make([]int, n1+1)
		n1++
		digits1[0] = 1
		if n1 < n2 {
			target := make([]int, n1)
			for i := 0; i < n1; i++ {
				target[i] = 9
			}
			out += subTotalWaviness(digits1, target)
		}
	}
	return out + subTotalWaviness(digits1, digits2)
}
