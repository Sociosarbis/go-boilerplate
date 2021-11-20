package lhs

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestLHS(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 3, 2, 2, 5, 2, 3, 7},
			5,
		},
		{
			[]int{1, 2, 3, 4},
			2,
		},
		{
			[]int{1, 1, 1, 1},
			0,
		},
	}

	for _, s := range suites {
		ret := findLHS(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
