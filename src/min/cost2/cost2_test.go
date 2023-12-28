package cost2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	x    int
	ret  int64
}

func TestCost(t *testing.T) {
	suites := []suite{
		{
			[]int{20, 1, 15},
			5,
			13,
		},
		{
			[]int{1, 2, 3},
			4,
			6,
		},
	}

	for _, s := range suites {
		ret := minCost(s.nums, s.x)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
