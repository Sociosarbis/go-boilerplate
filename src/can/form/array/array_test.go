package array

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr    []int
	pieces [][]int
	ret    bool
}

func TestArray(t *testing.T) {
	suites := []suite{
		{
			[]int{15, 88},
			[][]int{{88}, {15}},
			true,
		},
		{
			[]int{49, 18, 16},
			[][]int{{16, 18, 49}},
			false,
		},
		{
			[]int{91, 4, 64, 78},
			[][]int{{78}, {4, 64}, {79}},
			true,
		},
	}

	for _, s := range suites {
		ret := canFormArray(s.arr, s.pieces)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
