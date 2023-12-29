package choco

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	prices []int
	money  int
	ret    int
}

func TestChoco(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 2},
			3,
			0,
		},
		{
			[]int{3, 2, 3},
			3,
			3,
		},
	}

	for _, s := range suites {
		ret := buyChoco(s.prices, s.money)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
