package possible

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	blocked [][]int
	source  []int
	target  []int
	ret     bool
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 1}, {1, 0}},
			[]int{0, 0},
			[]int{0, 2},
			false,
		},
		{
			[][]int{},
			[]int{0, 0},
			[]int{999999, 999999},
			true,
		},
	}

	for _, s := range suites {
		ret := isEscapePossible(s.blocked, s.source, s.target)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
