package choose

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	groups [][]int
	nums   []int
	ret    bool
}

func TestChoose(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, -1, -1}, {3, -2, 0}},
			[]int{1, -1, 0, 1, -1, -1, 3, -2, 0},
			true,
		},
		{
			[][]int{{10, -2}, {1, 2, 3, 4}},
			[]int{1, 2, 3, 4, 10, -2},
			false,
		},
		{
			[][]int{{1, 2, 3}, {3, 4}},
			[]int{7, 7, 1, 2, 3, 4, 7, 7},
			false,
		},
	}
	for _, s := range suites {
		ret := canChoose(s.groups, s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
