package time

type item struct {
	i int
	t int
}

func addItem(queue *[]item, m item) {
	l := 0
	r := len(*queue) - 1
	for l <= r {
		mid := (l + r) / 2
		if (*queue)[mid].t > m.t {
			l = mid + 1
		} else if (*queue)[mid].t < m.t {
			if mid > 0 && (*queue)[mid-1].t < m.t {
				r = mid - 1
			} else {
				l = mid
				break
			}
		} else {
			l = mid
			break
		}
	}
	if l >= len(*queue) {
		*queue = append(*queue, m)
	} else {
		*queue = append(*queue, item{})
		copy((*queue)[l+1:], (*queue)[l:])
		(*queue)[l] = m
	}
}

func minimumTime(n int, relations [][]int, time []int) int {
	deps := make([]int, n+1)
	nexts := make([][]int, n+1)
	for _, r := range relations {
		deps[r[1]]++
		nexts[r[0]] = append(nexts[r[0]], r[1])
	}
	readys := []item{}
	ret := 0
	for i := 1; i <= n; i++ {
		if deps[i] == 0 {
			addItem(&readys, item{
				i,
				time[i-1],
			})
		}
	}
	for len(readys) != 0 {
		backIndex := len(readys) - 1
		first := readys[backIndex]
		ret = first.t
		readys = readys[:backIndex]
		for _, n := range nexts[first.i] {
			deps[n]--
			if deps[n] == 0 {
				addItem(&readys, item{
					n,
					ret + time[n-1],
				})
			}
		}
	}
	return ret
}
