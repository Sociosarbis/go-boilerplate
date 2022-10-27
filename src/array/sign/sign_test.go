package sign

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestSign(t *testing.T) {
	suites := []suite{
		{
			[]int{-1, -2, -3, -4, 3, 2, 1},
			1,
		},
		{
			[]int{1, 5, 0, 2, -3},
			0,
		},
		{
			[]int{-1, 1, -1, 1, -1},
			-1,
		},
	}

	for _, s := range suites {
		ret := arraySign(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
