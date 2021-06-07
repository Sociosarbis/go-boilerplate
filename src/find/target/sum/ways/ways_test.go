package ways

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums   []int
	target int
	ret    int
}

func TestWays(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 1, 1, 1, 1},
			3,
			5,
		},
		{
			[]int{1},
			1,
			1,
		},
		{
			[]int{2, 107, 109, 113, 127, 131, 137, 3, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 47, 53},
			1000,
			0,
		},
		{
			[]int{1, 0},
			1,
			2,
		},
	}

	for _, s := range suites {
		ret := findTargetSumWays(s.nums, s.target)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
