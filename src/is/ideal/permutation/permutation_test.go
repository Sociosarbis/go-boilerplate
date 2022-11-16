package permutation

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  bool
}

func TestPermutation(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 0, 2},
			true,
		},
		{
			[]int{1, 2, 0},
			false,
		},
	}

	for _, s := range suites {
		ret := isIdealPermutation(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
