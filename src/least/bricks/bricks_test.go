package bricks

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	wall [][]int
	ret  int
}

func TestBricks(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}},
			2,
		},
		{
			[][]int{{1}, {1}, {1}},
			3,
		},
	}

	for _, s := range suites {
		ret := leastBricks(s.wall)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
