package ancestors

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	ins := make([]int, n)
	g := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		ins[b]++
	}

	temp := make([]map[int]struct{}, n)

	for i := 0; i < n; i++ {
		temp[i] = map[int]struct{}{}
	}

	queue := []int{}

	for i, count := range ins {
		if count == 0 {
			queue = append(queue, i)
		}
	}

	ret := make([][]int, n)
	n = len(queue)
	for n != 0 {
		for i := 0; i < n; i++ {
			p := queue[i]
			for _, c := range g[p] {
				ins[c]--
				if ins[c] == 0 {
					queue = append(queue, c)
				}
				temp[c][p] = struct{}{}
				for k := range temp[p] {
					if _, ok := temp[c][k]; !ok {
						temp[c][k] = struct{}{}
					}
				}
			}
		}
		queue = queue[n:]
		n = len(queue)
	}

	for i := range ret {
		ret[i] = make([]int, 0, len(temp[i]))
		for k := range temp[i] {
			ret[i] = append(ret[i], k)
		}
		sort.Ints(ret[i])
	}

	return ret
}
