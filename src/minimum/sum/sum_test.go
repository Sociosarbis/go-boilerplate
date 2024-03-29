package sum

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestSum(t *testing.T) {
	suites := []suite{
		{
			[]int{8, 6, 1, 5, 3},
			9,
		},
		{
			[]int{5, 4, 8, 7, 10, 2},
			13,
		},
		{
			[]int{6, 5, 4, 3, 4, 5},
			-1,
		},
	}

	for _, s := range suites {
		ret := minimumSum(s.nums)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
