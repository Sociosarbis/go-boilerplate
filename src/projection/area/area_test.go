package area

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  int
}

func TestArea(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2}, {3, 4}},
			17,
		},
		{
			[][]int{{2}},
			5,
		},
		{
			[][]int{{1, 0}, {0, 2}},
			8,
		},
	}

	for _, s := range suites {
		ret := projectionArea(s.grid)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
