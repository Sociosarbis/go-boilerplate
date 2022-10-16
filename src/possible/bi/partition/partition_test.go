package partition

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n        int
	dislikes [][]int
	ret      bool
}

func TestPartition(t *testing.T) {
	suites := []suite{
		{
			4,
			[][]int{{1, 2}, {1, 3}, {2, 4}},
			true,
		},
		{
			3,
			[][]int{{1, 2}, {1, 3}, {2, 3}},
			false,
		},
		{
			5,
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}},
			false,
		},
	}

	for _, s := range suites {
		ret := possibleBipartition(s.n, s.dislikes)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
