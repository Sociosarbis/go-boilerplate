package target

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type Suite struct {
	matrix [][]int
	target int
	ret    int
}

func TestNumSubmatrixSumTarget(t *testing.T) {
	suites := []Suite{
		{
			[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}},
			0,
			4,
		},
		{
			[][]int{{1, -1}, {-1, 1}},
			0,
			5,
		},
		{
			[][]int{{904}},
			0,
			0,
		},
	}

	for _, su := range suites {
		ret := numSubmatrixSumTarget(su.matrix, su.target)
		if ret != su.ret {
			t.Errorf(constant.ErrTemplateStr, su.ret, numSubmatrixSumTarget(su.matrix, su.target))
		}
	}
}
