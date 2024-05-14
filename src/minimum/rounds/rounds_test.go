package rounds

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	tasks []int
	ret   int
}

func TestRounds(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4},
			4,
		},
		{
			[]int{2, 3, 3},
			-1,
		},
	}

	for _, s := range suites {
		ret := minimumRounds(s.tasks)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
