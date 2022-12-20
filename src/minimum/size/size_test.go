package size

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums          []int
	maxOperations int
	ret           int
}

func TestSize(t *testing.T) {
	suites := []suite{
		{
			[]int{9},
			2,
			3,
		},
		{
			[]int{2, 4, 8, 2},
			4,
			2,
		},
	}

	for _, s := range suites {
		ret := minimumSize(s.nums, s.maxOperations)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
