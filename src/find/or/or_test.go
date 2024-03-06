package or

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestOr(t *testing.T) {
	suites := []suite{
		{
			[]int{7, 12, 9, 8, 9, 15},
			4,
			9,
		},
		{
			[]int{2, 12, 1, 11, 4, 5},
			6,
			0,
		},
		{
			[]int{10, 8, 5, 9, 11, 6, 8},
			1,
			15,
		},
	}

	for _, s := range suites {
		ret := findKOr(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
