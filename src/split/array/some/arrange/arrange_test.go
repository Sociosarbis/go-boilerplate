package arrange

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  bool
}

func TestArrange(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			true,
		},
		{
			[]int{3, 1},
			false,
		},
	}

	for _, s := range suites {
		ret := splitArraySameAverage(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
