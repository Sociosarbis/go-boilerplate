package transform

import (
	"sort"
)

type node struct {
	next []*node
	row  int
	col  int
	v    int
	deps int
	g    int
}

type group struct {
	deps    int
	members []*node
}

type nodes []*node

func (this nodes) Len() int {
	return len(this)
}

func (this nodes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this nodes) Less(i, j int) bool {
	return this[i].v < this[j].v
}

func matrixRankTransform(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	nodeGraph := make([][]*node, m)
	groups := map[int]*group{}
	gid := 1
	for i := range nodeGraph {
		nodeGraph[i] = make([]*node, n)
		for j := range nodeGraph[i] {
			nodeGraph[i][j] = &node{
				row: i,
				col: j,
				v:   matrix[i][j],
			}
		}
	}

	for i := 0; i < m; i++ {
		arr := make([]*node, n)
		copy(arr, nodeGraph[i])
		sort.Sort(nodes(arr))
		j := 0
		prevRange := []int{0, 0}
		for j < n {
			prevJ := j
			groupId := 0
			for j < n && arr[j].v == arr[prevJ].v {
				arr[j].deps += prevRange[1] - prevRange[0]
				if groupId == 0 && arr[j].g != 0 {
					groupId = arr[j].g
				}
				for k := prevRange[0]; k < prevRange[1]; k++ {
					arr[k].next = append(arr[k].next, arr[j])
				}
				j++
			}
			if prevJ+1 != j {
				if groupId == 0 {
					groupId = gid
					groups[groupId] = &group{
						deps:    0,
						members: []*node{},
					}
					gid++
				}
				group := groups[groupId]
				for k := prevJ; k < j; k++ {
					if arr[k].g != groupId {
						if arr[k].g == 0 {
							group.members = append(group.members, arr[k])
							arr[k].g = groupId
						} else {
							oldGroupId := arr[k].g
							if other, ok := groups[oldGroupId]; ok {
								group.members = append(group.members, other.members...)
								for _, m := range other.members {
									m.g = groupId
								}
							}
							delete(groups, oldGroupId)
						}
					}
				}
			}
			prevRange[0], prevRange[1] = prevJ, j
		}
	}
	for i := 0; i < n; i++ {
		arr := make([]*node, m)
		for j := 0; j < m; j++ {
			arr[j] = nodeGraph[j][i]
		}
		sort.Sort(nodes(arr))
		j := 0
		prevRange := []int{0, 0}
		for j < m {
			prevJ := j
			groupId := 0
			for j < m && arr[j].v == arr[prevJ].v {
				arr[j].deps += prevRange[1] - prevRange[0]
				if groupId == 0 && arr[j].g != 0 {
					groupId = arr[j].g
				}
				for k := prevRange[0]; k < prevRange[1]; k++ {
					arr[k].next = append(arr[k].next, arr[j])
				}
				j++
			}
			if prevJ+1 != j {
				if groupId == 0 {
					groupId = gid
					groups[groupId] = &group{}
					gid++
				}
				group := groups[groupId]
				for k := prevJ; k < j; k++ {
					if arr[k].g != groupId {
						if arr[k].g == 0 {
							group.members = append(group.members, arr[k])
							arr[k].g = groupId
						} else {
							oldGroupId := arr[k].g
							if other, ok := groups[oldGroupId]; ok {
								group.members = append(group.members, other.members...)
								for _, m := range other.members {
									m.g = groupId
								}
							}
							delete(groups, oldGroupId)
						}
					}
				}
			}
			prevRange[0], prevRange[1] = prevJ, j
		}
	}

	top := []*node{}

	for _, group := range groups {
		for _, m := range group.members {
			group.deps += m.deps
		}
	}

	for _, row := range nodeGraph {
		for _, cell := range row {
			if cell.deps == 0 && (cell.g == 0 || groups[cell.g].deps == 0) {
				top = append(top, cell)
			}
		}
	}

	i := 1
	for len(top) != 0 {
		count := len(top)
		for j := 0; j < count; j++ {
			matrix[top[j].row][top[j].col] = i
			for _, next := range top[j].next {
				next.deps--
				if next.g == 0 {
					if next.deps == 0 {
						top = append(top, next)
					}
				} else {
					if group, ok := groups[next.g]; ok {
						group.deps--
					}
					if groups[next.g].deps == 0 {
						top = append(top, groups[next.g].members...)
					}
				}
			}
		}
		top = top[count:]
		i++
	}

	return matrix
}
