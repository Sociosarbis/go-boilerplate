package equal

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	target []int
	arr    []int
	ret    bool
}

func TestEqual(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3, 4},
			[]int{2, 4, 1, 3},
			true,
		},
		{
			[]int{7},
			[]int{7},
			true,
		},
		{
			[]int{3, 7, 9},
			[]int{3, 7, 11},
			false,
		},
	}
	for _, s := range suites {
		ret := canBeEqual(s.target, s.arr)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
