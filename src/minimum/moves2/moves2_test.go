package moves2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums       []int
	k          int
	maxChanges int
	ret        int64
}

func TestMoves(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 0, 1, 0, 1},
			3,
			0,
			4,
		},
	}

	for _, s := range suites {
		ret := minimumMoves(s.nums, s.k, s.maxChanges)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
