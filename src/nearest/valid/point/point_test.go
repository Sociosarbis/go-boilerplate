package point

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	x      int
	y      int
	points [][]int
	ret    int
}

func TestPoint(t *testing.T) {
	suites := []suite{
		{
			3,
			4,
			[][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}},
			2,
		},
		{
			3,
			4,
			[][]int{{3, 4}},
			0,
		},
		{
			3,
			4,
			[][]int{{2, 3}},
			-1,
		},
	}

	for _, s := range suites {
		ret := nearestValidPoint(s.x, s.y, s.points)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
