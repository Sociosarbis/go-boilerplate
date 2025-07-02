package count2

const mask int = 1e9 + 7

func possibleStringCount(word string, k int) int {
	n := len(word)
	if n < k {
		return 0
	}
	dp := [2][]int{
		make([]int, k),
		make([]int, k),
	}
	dp[1][0] = 1
	var total int64 = 1
	var i, index, lastK, count int
	last := 1
	for i < n {
		j := i
		for ; j < n; j++ {
			if word[i] != word[j] {
				break
			}
		}
		c := j - i
		total = (total * int64(c)) % int64(mask)
		if count < k {
			count++
			prevIndex := 1 - index
			kk := k
			if kk > j+1 {
				kk = j + 1
			}
			var ii int
			if ii < index-2 {
				ii = index - 2
			}
			for ; ii < kk; ii++ {
				dp[index][ii] = 0
				if ii < index+1 {
					continue
				}
				start := ii - c
				if start < index {
					start = index
				}
				end := ii
				if i == 0 {
					end++
				}
				var sum int
				if start < end {
					if end > 0 {
						sum = dp[prevIndex][end-1]
					}
					if sum == 0 && end >= lastK {
						sum = last
					}
					if start > 0 {
						sum -= dp[prevIndex][start-1]
						if sum < 0 {
							sum += mask
						}
					}
				}
				if ii > 0 {
					dp[index][ii] = (sum + dp[index][ii-1]) % mask
				} else {
					dp[index][ii] = sum
				}
			}
			last = dp[index][kk-1]
			lastK = kk
		}
		index = (index + 1) & 1
		i = j
	}
	index = 1 - index
	if count >= k {
		dp[index][k-1] = 0
	}
	out := int(total) - dp[index][k-1]
	if out < 0 {
		out += mask
	}
	return out
}
