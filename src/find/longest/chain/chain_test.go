package chain

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	pairs [][]int
	ret   int
}

func TestChain(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}},
			2,
		},
		{
			[][]int{{1, 2}, {7, 8}, {4, 5}},
			3,
		},
	}

	for _, s := range suites {
		ret := findLongestChain(s.pairs)
		for ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
