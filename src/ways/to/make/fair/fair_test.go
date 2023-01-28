package fair

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestFair(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 1, 6, 4},
			1,
		},
		{
			[]int{1, 1, 1},
			3,
		},
		{
			[]int{1, 2, 3},
			0,
		},
	}

	for _, s := range suites {
		ret := waysToMakeFair(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
