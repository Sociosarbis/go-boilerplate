package division

import "fmt"

type division struct {
	val  float64
	text string
}

func concat(s ...string) string {
	length := 0
	for _, item := range s {
		length += len(item)
	}
	b := make([]byte, length)
	i := 0
	for _, item := range s {
		i += copy(b[i:], item)
	}
	return string(b)
}

func optimalDivision(nums []int) string {
	dp := [10][10][2]*division{}
	for i := 1; i <= len(nums); i += 1 {
		for j := 0; j <= len(nums)-i; j += 1 {
			if i == 1 {
				dp[i-1][j][0] = &division{
					float64(nums[j]),
					fmt.Sprint(nums[j]),
				}
				dp[i-1][j][1] = dp[i-1][j][0]
			} else {
				for k := 1; k < i; k += 1 {
					val := dp[k-1][j][0].val / dp[i-k-1][j+k][1].val
					if dp[i-1][j][0] == nil || dp[i-1][j][0].val > val {
						var text string
						if i-k == 1 {
							text = concat(dp[k-1][j][0].text, "/", dp[i-k-1][j+k][1].text)
						} else {
							text = concat(dp[k-1][j][0].text, "/(", dp[i-k-1][j+k][1].text, ")")
						}
						dp[i-1][j][0] = &division{
							val,
							text,
						}
					}
					val = dp[k-1][j][1].val / dp[i-k-1][j+k][0].val
					if dp[i-1][j][1] == nil || dp[i-1][j][1].val < val {
						var text string
						if i-k == 1 {
							text = concat(dp[k-1][j][1].text, "/", dp[i-k-1][j+k][0].text)
						} else {
							text = concat(dp[k-1][j][1].text, "/(", dp[i-k-1][j+k][0].text, ")")
						}
						dp[i-1][j][1] = &division{
							val,
							text,
						}
					}
				}
			}
		}
	}
	return dp[len(nums)-1][0][1].text
}
