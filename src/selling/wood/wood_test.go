package wood

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	m      int
	n      int
	prices [][]int
	ret    int64
}

func TestWood(t *testing.T) {
	suites := []suite{
		{
			3,
			5,
			[][]int{{1, 4, 2}, {2, 2, 7}, {2, 1, 3}},
			19,
		},
		{
			4,
			6,
			[][]int{{3, 2, 10}, {1, 4, 2}, {4, 1, 3}},
			32,
		},
	}

	for _, s := range suites {
		ret := sellingWood(s.m, s.n, s.prices)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
