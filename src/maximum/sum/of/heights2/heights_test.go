package heights2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	maxHeights []int
	ret        int64
}

func TestHeights(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 3, 4, 1, 1},
			13,
		},
		{
			[]int{6, 5, 3, 9, 2, 7},
			22,
		},
		{
			[]int{3, 2, 5, 5, 2, 3},
			18,
		},
	}

	for _, s := range suites {
		ret := maximumSumOfHeights(s.maxHeights)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
