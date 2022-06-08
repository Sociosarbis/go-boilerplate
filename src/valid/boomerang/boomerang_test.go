package boomerang

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	points [][]int
	ret    bool
}

func TestBoomerang(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 1}, {2, 3}, {3, 2}},
			true,
		},
		{
			[][]int{{1, 1}, {2, 2}, {3, 3}},
			false,
		},
	}

	for _, s := range suites {
		ret := isBoomerang(s.points)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
