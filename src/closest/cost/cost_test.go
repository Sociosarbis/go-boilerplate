package cost

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	baseCosts    []int
	toppingCosts []int
	target       int
	ret          int
}

func TestCost(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 7},
			[]int{3, 4},
			10,
			10,
		},
		{
			[]int{2, 3},
			[]int{4, 5, 100},
			18,
			17,
		},
		{
			[]int{3, 10},
			[]int{2, 5},
			9,
			8,
		},
	}

	for _, s := range suites {
		ret := closestCost(s.baseCosts, s.toppingCosts, s.target)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
