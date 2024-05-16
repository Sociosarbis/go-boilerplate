package weeks

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	milestones []int
	ret        int64
}

func TestWeeks(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3},
			6,
		},
		{
			[]int{5, 2, 1},
			7,
		},
	}

	for _, s := range suites {
		ret := numberOfWeeks(s.milestones)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
