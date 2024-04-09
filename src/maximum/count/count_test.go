package count

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestCount(t *testing.T) {
	suites := []suite{
		{
			[]int{-2, -1, -1, 1, 2, 3},
			3,
		},
		{
			[]int{-3, -2, -1, 0, 0, 1, 2},
			3,
		},
		{
			[]int{5, 20, 66, 1314},
			4,
		},
	}

	for _, s := range suites {
		ret := maximumCount(s.nums)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
