package sign

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n     int
	mines [][]int
	ret   int
}

func TestSign(t *testing.T) {
	suites := []suite{
		{
			5,
			[][]int{{
				4,
				2,
			}},
			2,
		},
		{
			1,
			[][]int{{0, 0}},
			0,
		},
	}

	for _, s := range suites {
		ret := orderOfLargestPlusSign(s.n, s.mines)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
