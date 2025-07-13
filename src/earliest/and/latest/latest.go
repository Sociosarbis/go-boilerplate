package latest

type empty struct{}

func derive(out *[]int, visited map[int]empty, mask, l, r, firstPlayer, secondPlayer int) {
	for l <= r {
		if mask&(1<<l) != 0 {
			break
		} else {
			l++
		}
	}
	for l <= r {
		if mask&(1<<r) != 0 {
			break
		} else {
			r--
		}
	}
	if l == r {
		if _, ok := visited[mask]; !ok {
			*out = append(*out, mask)
			visited[mask] = empty{}
		}
	} else if l < r {
		if !((l+1 == secondPlayer && r+1 != firstPlayer) || l+1 == firstPlayer) {
			derive(out, visited, mask-(1<<l), l+1, r-1, firstPlayer, secondPlayer)
		}
		if !((r+1 == secondPlayer && l+1 != firstPlayer) || r+1 == firstPlayer) {
			derive(out, visited, mask-(1<<r), l+1, r-1, firstPlayer, secondPlayer)
		}
	} else {
		if _, ok := visited[mask]; !ok {
			*out = append(*out, mask)
			visited[mask] = empty{}
		}
	}
}

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	mask := (1 << n) - 1
	min, max := n, 1
	bfs := []int{mask}
	visited := map[int]empty{}
	checkMask := (1 << (firstPlayer - 1)) | (1 << (secondPlayer - 1))
	l := len(bfs)
	var count int
	for l != 0 {
		var i int
		count++
		for ; i < l; i++ {
			derive(&bfs, visited, bfs[i], 0, n-1, firstPlayer, secondPlayer)
		}
		l = 0
		for ; i < len(bfs); i++ {
			if bfs[i]&checkMask != checkMask {
				if count < min {
					min = count
				}
				if count > max {
					max = count
				}
			} else {
				bfs[l] = bfs[i]
				l++
			}
		}
		bfs = bfs[:l]
	}
	return []int{min, max}
}
