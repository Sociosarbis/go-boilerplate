package doubled

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret bool
}

func TestDoubled(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 1, 3, 6},
			false,
		},
		{
			[]int{2, 1, 2, 6},
			false,
		},
		{
			[]int{4, -2, 2, -4},
			true,
		},
		{
			[]int{2, 4, 0, 0, 8, 1},
			true,
		},
	}

	for _, s := range suites {
		ret := canReorderDoubled(s.arr)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, ret, s.ret)
		}
	}
}
