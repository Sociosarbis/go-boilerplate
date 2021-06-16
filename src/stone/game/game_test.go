package game

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	piles []int
	ret   bool
}

func TestStoneGame(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 3, 4, 5},
			true,
		},
	}

	for _, s := range suites {
		ret := stoneGame(s.piles)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
