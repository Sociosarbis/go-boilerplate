package vii

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	stones []int
	ret    int
}

func TestVII(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 3, 1, 4, 2},
			6,
		},
		{
			[]int{7, 90, 5, 1, 100, 10, 10, 2},
			122,
		},
	}

	for _, s := range suites {
		ret := stoneGameVII(s.stones)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
