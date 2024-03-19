package score2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestScore(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 4, 3, 7, 4, 5},
			3,
			15,
		},
		{
			[]int{5, 5, 4, 5, 4, 1, 1, 1},
			0,
			20,
		},
	}

	for _, s := range suites {
		ret := maximumScore(s.nums, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
