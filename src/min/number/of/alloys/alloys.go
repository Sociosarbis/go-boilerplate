package alloys

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	r := stock[0] + budget/cost[0]
	for i := 1; i < n; i++ {
		temp := stock[i] + budget/cost[i]
		if r < temp {
			r = temp
		}
	}
	l := 0
	ret := 0
loop:
	for l <= r {
		mid := (l + r) >> 1
		for i := 0; i < k; i++ {
			c := composition[i]
			t1 := 0
			for j := 0; j < n; j++ {
				t2 := c[j]*mid - stock[j]
				if t2 > 0 {
					t1 += t2 * cost[j]
				}
			}
			if t1 <= budget {
				ret = mid
				l = mid + 1
				continue loop
			}
		}
		r = mid - 1
	}
	return ret
}
