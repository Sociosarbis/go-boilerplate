package disjoint

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestDisjoint(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 0, 3, 8, 6},
			3,
		},
		{
			[]int{1, 1, 1, 0, 6, 12},
			4,
		},
	}

	for _, s := range suites {
		ret := partitionDisjoint(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
