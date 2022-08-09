package value

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestValue(t *testing.T) {
	suites := []suite{
		{
			[]int{-3, 2, -3, 4, 2},
			5,
		},
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{1, -2, -3},
			5,
		},
	}

	for _, s := range suites {
		ret := minStartValue(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
