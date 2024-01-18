package removal

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	beans []int
	ret   int64
}

func TestRemoval(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 1, 6, 5},
			4,
		},
		{
			[]int{2, 10, 3, 2},
			7,
		},
	}

	for _, s := range suites {
		ret := minimumRemoval(s.beans)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
