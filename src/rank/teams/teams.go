package teams

import "sort"

type empty struct{}

func rankTeams(votes []string) string {
	chars := make(map[byte]empty, 26)
	counter := [26][26]int{}
	for _, vote := range votes {
		for j, v := range vote {
			counter[v-65][j]++
			if _, ok := chars[byte(v)]; !ok {
				chars[byte(v)] = empty{}
			}
		}
	}
	n := len(chars)
	ret := make([]byte, 0, n)

	for k := range chars {
		ret = append(ret, k)
	}

	sort.Slice(ret, func(i, j int) bool {
		a, b := ret[i]-65, ret[j]-65
		for k := 0; k < n; k++ {
			if counter[a][k] > counter[b][k] {
				return true
			} else if counter[a][k] < counter[b][k] {
				return false
			}
		}
		return a < b
	})

	return string(ret)
}
