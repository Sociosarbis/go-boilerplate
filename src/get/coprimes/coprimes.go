package coprimes

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func getCoprimes(nums []int, edges [][]int) []int {
	g := map[int][]int{}
	for _, edge := range edges {
		for _, item := range [][2]int{{edge[0], edge[1]}, {edge[1], edge[0]}} {
			if _, ok := g[item[0]]; !ok {
				g[item[0]] = []int{item[1]}
			} else {
				g[item[0]] = append(g[item[0]], item[1])
			}
		}
	}
	closestParent := make([][2]int, 51)
	for i := 1; i <= 50; i++ {
		closestParent[i][0], closestParent[i][1] = -1, -1
	}
	n := len(nums)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = -1
	}
	dfs(g, nums, 0, -1, 0, closestParent, ret)
	return ret
}

func dfs(g map[int][]int, nums []int, c, p, depth int, closestParent [][2]int, out []int) {
	num := nums[c]
	temp := -1
	for i := 1; i <= 50; i++ {
		if closestParent[i][0] > temp {
			if gcd(num, i) == 1 {
				temp, out[c] = closestParent[i][0], closestParent[i][1]
			}
		}
	}
	oldDepth := closestParent[num]
	closestParent[num] = [2]int{depth, c}
	for _, next := range g[c] {
		if next != p {
			dfs(g, nums, next, c, depth+1, closestParent, out)
		}
	}
	closestParent[num] = oldDepth
}
