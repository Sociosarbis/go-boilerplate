package island

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  int
}

func TestIsland(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 0}, {0, 1}},
			3,
		},
		{
			[][]int{{1, 1}, {1, 0}},
			4,
		},
		{
			[][]int{{1, 1}, {1, 1}},
			4,
		},
	}

	for _, s := range suites {
		ret := largestIsland(s.grid)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
