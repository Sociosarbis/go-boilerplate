package ways

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	ranges [][]int
	ret    int
}

func TestWays(t *testing.T) {
	suites := []suite{
		{
			[][]int{{6, 10}, {5, 15}},
			2,
		},
		{
			[][]int{{1, 3}, {10, 20}, {2, 5}, {4, 8}},
			4,
		},
	}

	for _, s := range suites {
		ret := countWays(s.ranges)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
