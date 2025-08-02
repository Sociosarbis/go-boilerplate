package cost4

import (
	"sort"
)

func minCost(basket1 []int, basket2 []int) int64 {
	n := len(basket1)
	counter1, counter2 := make(map[int]int, n), make(map[int]int, n)
	min := basket1[0]
	for _, it := range basket1 {
		if it < min {
			min = it
		}
		if c, ok := counter1[it]; ok {
			counter1[it] = c + 1
		} else {
			counter1[it] = 1
		}
	}
	for _, it := range basket2 {
		if it < min {
			min = it
		}
		if c, ok := counter2[it]; ok {
			counter2[it] = c + 1
		} else {
			counter2[it] = 1
		}
	}
	more, less := make([]int, 0, n), make([]int, 0, n)
	for k, v := range counter1 {
		if c, ok := counter2[k]; ok {
			if (c+v)&1 == 1 {
				return -1
			}
			if v > c {
				more = append(more, k)
			}
		} else {
			if v&1 == 1 {
				return -1
			}
			more = append(more, k)
		}
	}
	for k, v := range counter2 {
		if c, ok := counter1[k]; ok {
			if (c+v)&1 == 1 {
				return -1
			}
			if v > c {
				less = append(less, k)
			}
		} else {
			if v&1 == 1 {
				return -1
			}
			less = append(less, k)
		}
	}
	sort.Ints(less)
	sort.Ints(more)
	var i, j int
	var out int64
	for i < len(less) && j < len(more) {
		if less[i] == min || (less[i] < more[j] && min*2 > less[i]) {
			out += int64(less[i])
			c := counter1[less[i]]
			counter1[less[i]] = c + 1
			c2 := counter2[less[i]]
			counter2[less[i]] = c2 - 1
			if c+1 == c2-1 {
				i++
			}
			c = counter1[more[len(more)-1]]
			counter1[more[len(more)-1]] = c - 1
			c2 = counter2[more[len(more)-1]]
			counter2[more[len(more)-1]] = c2 + 1
			if c-1 == c2+1 {
				more = more[:len(more)-1]
			}
		} else if more[j] == min || (more[j] < less[i] && min*2 > more[j]) {
			out += int64(more[j])
			c := counter1[more[j]]
			counter1[more[j]] = c - 1
			c2 := counter2[more[j]]
			counter2[more[j]] = c2 + 1
			if c-1 == c2+1 {
				j++
			}
			c = counter1[less[len(less)-1]]
			counter1[less[len(less)-1]] = c + 1
			c2 = counter2[less[len(less)-1]]
			counter2[less[len(less)-1]] = c2 - 1
			if c+1 == c2-1 {
				less = less[:len(less)-1]
			}
		} else {
			for ; i < len(less); i++ {
				c1, c2 := counter1[less[i]], counter2[less[i]]
				out += int64(c2-c1) / 2 * int64(min)
			}
			for ; j < len(more); j++ {
				c1, c2 := counter1[more[j]], counter2[more[j]]
				out += int64(c1-c2) / 2 * int64(min)
			}
		}
	}
	return out
}
