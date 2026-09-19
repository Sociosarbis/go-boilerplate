package substrings

import "sort"

type dm struct {
	indices []int
	length  int
	end     int
}

func maxNumOfSubstrings(s string) []string {
	ranges := make([][2]int, 26)
	for i, c := range s {
		index := c - 'a'
		if ranges[index][0] == 0 && ranges[index][1] == 0 {
			ranges[index][0], ranges[index][1] = i, i+1
		} else {
			ranges[index][1] = i + 1
		}
	}
	items := make([][2]int, 0, 26)
loop:
	for _, r := range ranges {
		if r[0] != 0 || r[1] != 0 {
			for i := r[0] + 1; i < r[1]; i++ {
				r2 := ranges[s[i]-'a']
				if r2[0] < r[0] {
					continue loop
				}
				if r2[1] > r[1] {
					r[1] = r2[1]
				}
			}
			items = append(items, r)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	in := len(items)
	dp := make([][]dm, in)
	for i := 0; i < in; i++ {
		dp[i] = make([]dm, in)
	}
	maxLength := len(s)
	var maxCount, maxI int
	for i, item := range items {
		start, length, end := item[0], item[1]-item[0], item[1]
		dp[i][0] = dm{
			indices: []int{i},
			length:  length,
			end:     end,
		}
		if maxCount == 0 && length < maxLength {
			maxLength = length
			maxI = i
		}
		for j := 1; j < in; j++ {
			for k := 0; k < i; k++ {
				if dp[k][j-1].length != 0 && dp[k][j-1].end <= start {
					if dp[i][j].length == 0 || dp[i][j].length > dp[k][j-1].length+length {
						dp[i][j].length = dp[k][j-1].length + length
						dp[i][j].end = end
						dp[i][j].indices = make([]int, j+1)
						copy(dp[i][j].indices, dp[k][j-1].indices)
						dp[i][j].indices[j] = i
						if j > maxCount || (j == maxCount && length < maxLength) {
							maxCount = j
							maxLength = length
							maxI = i
						}
					}
				}
			}
		}
	}
	indices := dp[maxI][maxCount].indices
	out := make([]string, 0, maxCount+1)
	for _, index := range indices {
		l, r := items[index][0], items[index][1]
		out = append(out, s[l:r])
	}
	return out
}
