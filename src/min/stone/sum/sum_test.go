package sum

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	piles []int
	k     int
	ret   int
}

func TestSum(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 4, 9},
			2,
			12,
		},
		{
			[]int{4, 3, 6, 7},
			3,
			12,
		},
	}

	for _, s := range suites {
		ret := minStoneSum(s.piles, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
