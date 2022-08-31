package sequences

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	pushed []int
	popped []int
	ret    bool
}

func TestSequences(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 1, 2},
			false,
		},
	}

	for _, s := range suites {
		ret := validateStackSequences(s.pushed, s.popped)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
