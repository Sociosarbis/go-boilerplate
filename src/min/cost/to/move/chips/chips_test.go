package chips

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	position []int
	ret      int
}

func TestChips(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3},
			1,
		},
		{
			[]int{2, 2, 2, 3, 3},
			2,
		},
		{
			[]int{1, 1000000000},
			1,
		},
	}
	for _, s := range suites {
		ret := minCostToMoveChips(s.position)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
