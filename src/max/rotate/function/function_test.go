package function

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestFunction(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 3, 2, 6},
			26,
		},
		{
			[]int{100},
			0,
		},
	}

	for _, s := range suites {
		ret := maxRotateFunction(s.nums)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
