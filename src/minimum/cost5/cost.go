package cost5

type node struct {
	next [26]*node
	id   int
}

func buildTrie(n *node, word string, id int) *node {
	cur := n
	for _, c := range word {
		if cur.next[c-'a'] == nil {
			cur.next[c-'a'] = &node{
				id: -1,
			}
		}
		cur = cur.next[c-'a']
	}
	cur.id = id
	return n
}

func min(a, b int64) int64 {
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	m := len(original)
	strToId := make(map[string]int, 2*m)
	var id int
	for _, s := range original {
		if _, ok := strToId[s]; !ok {
			strToId[s] = id
			id++
		}
	}
	for _, s := range changed {
		if _, ok := strToId[s]; !ok {
			strToId[s] = id
			id++
		}
	}
	dist := make([][]int64, id)
	for i := range dist {
		dist[i] = make([]int64, id)
		for j := 0; j < id; j++ {
			dist[i][j] = -1
		}
		dist[i][i] = 0
	}
	for j, s1 := range original {
		dist[strToId[s1]][strToId[changed[j]]] = min(dist[strToId[s1]][strToId[changed[j]]], int64(cost[j]))
	}
	// Floyd Warshall
	for k := 0; k < id; k++ {
		for i := 0; i < id; i++ {
			for j := 0; j < id; j++ {
				if dist[i][k] >= 0 && dist[k][j] >= 0 {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}
	r1 := &node{
		id: -1,
	}
	r2 := &node{
		id: -1,
	}
	for _, s := range original {
		r1 = buildTrie(r1, s, strToId[s])
	}
	for _, s := range changed {
		r2 = buildTrie(r2, s, strToId[s])
	}
	n := len(source) + 1
	dp := make([]int64, n)
	for i := n - 2; i >= 0; i-- {
		dp[i] = -1
	}
	for i := len(source) - 1; i >= 0; i-- {
		if source[i] == target[i] {
			dp[i] = dp[i+1]
		}
		c1 := r1
		c2 := r2
		for j := i + 1; j <= len(source); j++ {
			c1 = c1.next[source[j-1]-'a']
			c2 = c2.next[target[j-1]-'a']
			if c1 == nil || c2 == nil {
				break
			}
			if c1.id != -1 && c2.id != -1 && dist[c1.id][c2.id] != -1 && dp[j] != -1 {
				dp[i] = min(dp[i], dist[c1.id][c2.id]+dp[j])
			}
		}
	}
	return dp[0]
}
