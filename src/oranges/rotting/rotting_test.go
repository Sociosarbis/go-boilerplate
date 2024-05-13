package rotting

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  int
}

func TestRotting(t *testing.T) {
	suites := []suite{
		{
			[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			4,
		},
		{
			[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			-1,
		},
		{
			[][]int{{0, 2}},
			0,
		},
	}

	for _, s := range suites {
		ret := orangesRotting(s.grid)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
