package winners

import "sort"

func findWinners(matches [][]int) [][]int {
	g := map[int]byte{}

	for _, m := range matches {
		a, b := m[0], m[1]
		if _, ok := g[a]; !ok {
			g[a] = 0
		}
		if s, ok := g[b]; ok {
			if s < 2 {
				g[b] = s + 1
			}
		} else {
			g[b] = 1
		}
	}

	var n int
	for _, s := range g {
		if s == 0 {
			n++
		}
	}
	g1 := make([]int, 0, n)
	for num, s := range g {
		if s == 0 {
			g1 = append(g1, num)
		}
	}

	sort.Ints(g1)

	n = 0

	for _, s := range g {
		if s == 1 {
			n++
		}
	}
	g2 := make([]int, 0, n)
	for num, s := range g {
		if s == 1 {
			g2 = append(g2, num)
		}
	}

	sort.Ints(g2)

	return [][]int{g1, g2}
}
