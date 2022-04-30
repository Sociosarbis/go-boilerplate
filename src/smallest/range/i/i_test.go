package i

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestI(t *testing.T) {
	suites := []suite{
		{
			[]int{1},
			0,
			0,
		},
		{
			[]int{0, 10},
			2,
			6,
		},
		{
			[]int{1, 3, 6},
			3,
			0,
		},
	}

	for _, s := range suites {
		ret := smallestRangeI(s.nums, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
