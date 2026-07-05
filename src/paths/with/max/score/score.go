package score

import "fmt"

const mod int = 1e9 + 7

type item struct {
	c int
	a int
	b int
}

func (self *item) disable() {
	self.a, self.b, self.c = 0, 0, 0
}

func (self item) add(other item) item {
	if other.c == 0 {
		return self
	}
	if self.a == other.a {
		if self.b == other.b {
			self.c = (self.c + other.c) % mod
		}
	}
	return self
}

func (self item) max(other item) item {
	if other.c == 0 {
		return self
	} else if self.c == 0 {
		return other
	}
	if self.a == other.a {
		if self.b >= other.b {
			return self
		}
		return other
	} else if self.a > other.a {
		return self
	}
	return other
}

func pathsWithMaxScore(board []string) []int {
	n := len(board)
	dp := [2][]item{}
	for i := 0; i < 2; i++ {
		dp[i] = make([]item, n)
	}

	dp[0][n-1] = item{
		c: 1,
	}

	for j := n - 2; j >= 0; j-- {
		if board[n-1][j] != 'X' {
			dp[0][j] = dp[0][j+1]
			if dp[0][j].c != 0 && board[n-1][j] >= '1' && board[n-1][j] <= '9' {
				b := dp[0][j].b + int(board[n-1][j]-'0')
				dp[0][j].b = b % mod
				dp[0][j].a += b / mod
			}
		}
	}

	for i := n - 2; i >= 0; i-- {
		index := (n - 1 - i) & 1
		prevIndex := 1 - index
		fmt.Println(dp[prevIndex])
		for j := n - 1; j >= 0; j-- {
			var max item
			if j+1 < n {
				max = dp[index][j+1]
			} else if i+1 < n {
				max = dp[prevIndex][j]
			}
			if board[i][j] != 'X' {
				if i+1 < n {
					max = max.max(dp[prevIndex][j])
					if j+1 < n {
						max = max.max(dp[prevIndex][j+1])
					}
				}
				if j+1 < n {
					max = max.max(dp[index][j+1])
				}
				max.c = 0
				if i+1 < n {
					max = max.add(dp[prevIndex][j])
					if j+1 < n {
						max = max.add(dp[prevIndex][j+1])
					}
				}
				if j+1 < n {
					max = max.add(dp[index][j+1])
				}
				if max.c != 0 && board[i][j] >= '1' && board[i][j] <= '9' {
					b := max.b + int(board[i][j]-'0')
					max.b = b % mod
					max.a += b / mod
				}
			} else {
				max.disable()
			}
			dp[index][j] = max
		}
	}
	index := (n - 1) & 1
	return []int{dp[index][0].b, dp[index][0].c}
}
