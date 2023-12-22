package removals

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestRemovals(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 3, 1},
			0,
		},
		{
			[]int{2, 1, 1, 5, 6, 2, 3, 1},
			3,
		},
	}

	for _, s := range suites {
		ret := minimumMountainRemovals(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
