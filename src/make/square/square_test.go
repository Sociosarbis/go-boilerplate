package square

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	matchsticks []int
	ret         bool
}

func TestSquare(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 1, 2, 2, 2},
			true,
		},
		{
			[]int{3, 3, 3, 3, 4},
			false,
		},
	}

	for _, s := range suites {
		ret := makesquare(s.matchsticks)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
