package pickup

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  int
}

func TestPickup(t *testing.T) {
	suites := []suite{
		{
			[][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}},
			24,
		},
		{
			[][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}},
			28,
		},
	}

	for _, s := range suites {
		ret := cherryPickup(s.grid)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
