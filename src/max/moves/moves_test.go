package moves

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  int
}

func TestMoves(t *testing.T) {
	suites := []suite{
		{
			[][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}},
			3,
		},
		{
			[][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}},
			0,
		},
	}

	for _, s := range suites {
		ret := maxMoves(s.grid)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
