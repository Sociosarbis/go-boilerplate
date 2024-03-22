package cells

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  int
}

func TestCells(t *testing.T) {
	suites := []suite{
		{
			[][]int{{3, 4, 2, 1}, {4, 2, 3, 1}, {2, 1, 0, 0}, {2, 4, 0, 0}},
			4,
		},
		{
			[][]int{{3, 4, 2, 1}, {4, 2, 1, 1}, {2, 1, 1, 0}, {3, 4, 1, 0}},
			3,
		},
		{
			[][]int{{2, 1, 0}, {1, 0, 0}},
			-1,
		},
	}

	for _, s := range suites {
		ret := minimumVisitedCells(s.grid)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
