package count

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	m   int
	n   int
	ops [][]int
	ret int
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			3,
			3,
			[][]int{{2, 2}, {3, 3}},
			4,
		},
		{
			3,
			3,
			[][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}},
			4,
		},
		{
			3,
			3,
			[][]int{},
			9,
		},
	}

	for _, s := range suites {
		ret := maxCount(s.m, s.n, s.ops)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
