package square

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	p1  []int
	p2  []int
	p3  []int
	p4  []int
	ret bool
}

func TestSquare(t *testing.T) {
	suites := []suite{
		{
			[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}, true,
		},
		{
			[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 12}, false,
		},
		{
			[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}, true,
		},
		{
			[]int{-3, 1},
			[]int{0, 0},
			[]int{0, 10},
			[]int{3, 1},
			false,
		},
	}

	for _, s := range suites {
		ret := validSquare(s.p1, s.p2, s.p3, s.p4)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
