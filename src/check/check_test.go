package check

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  bool
}

func TestCheck(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 4, 5, 1, 2},
			true,
		},
		{
			[]int{2, 1, 3, 4},
			false,
		},
		{
			[]int{1, 2, 3},
			true,
		},
	}

	for _, s := range suites {
		ret := check(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
