package time

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	tasks [][]int
	ret   int
}

func TestTime(t *testing.T) {
	suites := []suite{
		{
			[][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}},
			2,
		},
		{
			[][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}},
			4,
		},
		{
			[][]int{{10, 16, 3}, {10, 20, 5}, {1, 12, 4}, {8, 11, 2}},
			6,
		},
		{
			[][]int{{1, 18, 5}, {3, 15, 1}},
			5,
		},
	}

	for _, s := range suites {
		ret := findMinimumTime(s.tasks)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
