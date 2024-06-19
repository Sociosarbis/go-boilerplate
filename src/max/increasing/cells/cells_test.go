package cells

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	mat [][]int
	ret int
}

func TestCells(t *testing.T) {
	suites := []suite{
		{
			[][]int{
				{4, 6}, {-3, -9}, {8, -4}, {5, 5}, {9, -2}, {3, -6}, {-7, 0},
			},
			7,
		},
	}

	for _, s := range suites {
		ret := maxIncreasingCells(s.mat)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
