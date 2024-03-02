package nodes2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n          int
	edges      [][]int
	restricted []int
	ret        int
}

func TestNodes(t *testing.T) {
	suites := []suite{
		{
			7,
			[][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}},
			[]int{4, 5},
			4,
		},
		{
			7,
			[][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}},
			[]int{4, 2, 1},
			3,
		},
	}

	for _, s := range suites {
		ret := reachableNodes(s.n, s.edges, s.restricted)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
