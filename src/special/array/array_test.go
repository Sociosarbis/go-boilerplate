package array

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestArray(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 5},
			2,
		},
		{
			[]int{0, 0},
			-1,
		},
		{
			[]int{0, 4, 3, 0, 4},
			3,
		},
	}

	for _, s := range suites {
		ret := specialArray(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
