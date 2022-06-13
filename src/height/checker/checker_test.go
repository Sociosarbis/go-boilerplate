package checker

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	heights []int
	ret     int
}

func TestChecker(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 1, 4, 2, 1, 3},
			3,
		},
		{
			[]int{5, 1, 2, 3, 4},
			5,
		},
		{
			[]int{1, 2, 3, 4, 5},
			0,
		},
	}

	for _, s := range suites {
		ret := heightChecker(s.heights)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
