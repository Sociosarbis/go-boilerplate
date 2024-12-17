package strings

type node struct {
	next map[byte]*node
}

func minValidStrings(words []string, target string) int {
	root := &node{
		next: map[byte]*node{},
	}
	for _, word := range words {
		n := len(word)
		cur := root
		for i := 0; i < n; i++ {
			if _, ok := cur.next[word[i]]; !ok {
				cur.next[word[i]] = &node{
					next: map[byte]*node{},
				}
			}
			cur = cur.next[word[i]]
		}
	}
	n := len(target)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = -1
	}
	for i := 0; i < n; i++ {
		if dp[i] == -1 {
			continue
		}
		cur := root
		temp := dp[i] + 1
		for j := i; j < n; j++ {
			if _, ok := cur.next[target[j]]; !ok {
				break
			}
			cur = cur.next[target[j]]
			if dp[j+1] == -1 || temp < dp[j+1] {
				if j+1 == n {
					return temp
				}
				dp[j+1] = temp
			}
		}
	}
	return dp[n]
}
