package value

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	parents := make([][]int, n)
	children := make([][]int, n)
	counter := make([]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		counter[a]++
		parents[b] = append(parents[b], a)
		children[a] = append(children[a], b)
	}
	var temp int
	var queue []int
	for i := 0; i < n; i++ {
		if counter[i] == 0 {
			queue = append(queue, i)
		}
	}
	l := len(queue)
	if l == 0 {
		return -1
	}
	dp := make([][26]int, n)
	temp += l
	var ret int
	for l != 0 {
		c := l
		for i := 0; i < c; i++ {
			a := queue[i]
			index := int(colors[a] - 'a')
			for j := 0; j < 26; j++ {
				var max int
				for _, child := range children[a] {
					if dp[child][j] > max {
						max = dp[child][j]
					}
				}
				if index == j {
					dp[a][j] = max + 1
					if ret <= max {
						ret = max + 1
					}
				} else {
					dp[a][j] = max
				}
			}
			for _, parent := range parents[a] {
				counter[parent]--
				if counter[parent] == 0 {
					queue = append(queue, parent)
				}
			}
		}
		queue = queue[c:]
		l = len(queue)
		temp += l
	}
	if temp < n {
		return -1
	}
	return ret
}
