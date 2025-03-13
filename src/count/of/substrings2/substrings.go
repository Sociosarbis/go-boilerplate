package substrings2

import "sort"

var mappings = [26]int{}

func init() {
	mappings[0], mappings['e'-'a'], mappings['i'-'a'], mappings['o'-'a'], mappings['u'-'a'] = 1, 2, 3, 4, 5
}

func getIndex(c byte) int {
	index := mappings[c-'a']
	if index > 0 {
		return index - 1
	} else {
		return 5
	}
}

func countOfSubstrings(word string, k int) int64 {
	n := len(word)
	prefix := make([][6]int, n+1)
	for i, c := range word {
		index := getIndex(byte(c))
		for j := 0; j < 6; j++ {
			if j == index {
				prefix[i+1][j] = prefix[i][j] + 1
			} else {
				prefix[i+1][j] = prefix[i][j]
			}
		}
	}
	var ret int64
	end := n - k - 5
	for i := 0; i <= end; i++ {
		// 第一个辅音大于k的位置
		index := sort.Search(n, func(j int) bool {
			if j < i {
				return false
			}
			return prefix[j+1][5]-prefix[i][5] > k
		})
		if index == n && prefix[index][5] < k {
			break
		}
		right := index
		// 第一个辅音等于k的位置
		left := sort.Search(n, func(j int) bool {
			if j < i {
				return false
			}
			return prefix[j+1][5]-prefix[i][5] >= k
		})
		index = sort.Search(n, func(j int) bool {
			if j < i {
				return false
			}
			for k := 0; k < 5; k++ {
				if prefix[j+1][k]-prefix[i][k] == 0 {
					return false
				}
			}
			return true
		})
		if index == n {
			break
		}
		if index > left {
			left = index
		}
		if left < right {
			ret += int64(right - left)
		}
	}
	return ret
}
