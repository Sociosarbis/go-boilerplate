package cost

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	costs      []int
	k          int
	candidates int
	ret        int64
}

func TestCost(t *testing.T) {
	suites := []suite{
		{
			[]int{17, 12, 10, 2, 7, 2, 11, 20, 8},
			3,
			4,
			11,
		},
		{
			[]int{1, 2, 4, 1},
			3,
			3,
			4,
		},
	}

	for _, s := range suites {
		ret := totalCost(s.costs, s.k, s.candidates)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
