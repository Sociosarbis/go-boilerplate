package walls

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	cost []int
	time []int
	ret  int
}

func TestWalls(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3, 2},
			[]int{1, 2, 3, 2},
			3,
		},
		{
			[]int{2, 3, 4, 2},
			[]int{1, 1, 1, 1},
			4,
		},
	}

	for _, s := range suites {
		ret := paintWalls(s.cost, s.time)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
