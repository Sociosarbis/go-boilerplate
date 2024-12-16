package room

import (
	"sort"
)

func closestRoom(rooms [][]int, queries [][]int) []int {
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i][1] < rooms[j][1]
	})
	n := len(queries)
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return queries[indices[i]][1] > queries[indices[j]][1]
	})
	m := len(rooms)
	options := make([]int, 0, m)
	ret := make([]int, n)
	i := len(rooms)
	for _, index := range indices {
		value := queries[index][1]
		j := sort.Search(i, func(i int) bool {
			return rooms[i][1] >= value
		})
		for k := j; k < i; k++ {
			value = rooms[k][0]
			index := sort.Search(len(options), func(i int) bool {
				return rooms[options[i]][0] > value
			})
			if index == len(options) {
				options = append(options, k)
			} else {
				options = append(options, 0)
				copy(options[index+1:], options[index:])
				options[index] = k
			}
		}
		if len(options) == 0 {
			ret[index] = -1
		} else {
			value = queries[index][0]
			ii := sort.Search(len(options), func(i int) bool {
				return rooms[options[i]][0] >= value
			})
			if ii < len(options) && (ii == 0 || rooms[options[ii]][0]-value < value-rooms[options[ii-1]][0]) {
				ret[index] = rooms[options[ii]][0]
			} else {
				ret[index] = rooms[options[ii-1]][0]
			}
		}
		i = j
	}
	return ret
}
