package change

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	coins  []int
	amount int
	ret    int
}

func TestChange(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			[]int{2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			0,
		},
	}

	for _, s := range suites {
		ret := coinChange(s.coins, s.amount)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
