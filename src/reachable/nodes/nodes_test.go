package nodes

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	edges    [][]int
	maxMoves int
	n        int
	ret      int
}

func TestNodes(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}},
			6,
			3,
			13,
		},
		{
			[][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}},
			10,
			4,
			23,
		},
		{
			[][]int{{1, 2, 4}, {1, 4, 5}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5}},
			17,
			5,
			1,
		},
	}
	for _, s := range suites {
		ret := reachableNodes(s.edges, s.maxMoves, s.n)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
