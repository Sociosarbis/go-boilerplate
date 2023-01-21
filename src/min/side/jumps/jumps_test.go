package jumps

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	obstacles []int
	ret       int
}

func TestJumps(t *testing.T) {
	suites := []suite{
		{
			[]int{0, 1, 2, 3, 0},
			2,
		},
		{
			[]int{0, 1, 1, 3, 3, 0},
			0,
		},
		{
			[]int{0, 2, 1, 0, 3, 0},
			2,
		},
	}

	for _, s := range suites {
		ret := minSideJumps(s.obstacles)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
