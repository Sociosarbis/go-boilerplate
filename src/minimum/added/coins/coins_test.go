package coins

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	coins  []int
	target int
	ret    int
}

func TestCoins(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 4, 10},
			19,
			2,
		},
		{
			[]int{1, 4, 10, 5, 7, 19},
			19,
			1,
		},
		{
			[]int{1, 1, 1},
			20,
			3,
		},
	}

	for _, s := range suites {
		ret := minimumAddedCoins(s.coins, s.target)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
