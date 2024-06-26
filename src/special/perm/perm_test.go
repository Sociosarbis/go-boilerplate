package perm

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestPerm(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 3, 6},
			2,
		}, {
			[]int{1, 4, 3},
			2,
		},
	}

	for _, s := range suites {
		ret := specialPerm(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
