package time

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n         int
	relations [][]int
	time      []int
	ret       int
}

func TestTime(t *testing.T) {
	suites := []suite{
		suite{
			3,
			[][]int{{1, 3}, {2, 3}},
			[]int{3, 2, 5},
			8,
		},
		suite{
			5,
			[][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}},
			[]int{1, 2, 3, 4, 5},
			12,
		},
	}
	for _, s := range suites {
		ret := minimumTime(s.n, s.relations, s.time)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
