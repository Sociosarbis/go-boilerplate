package score

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestScore(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{3, 4, 6, 8},
			11,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			14,
		},
	}

	for _, s := range suites {
		ret := maxScore(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
