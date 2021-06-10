package change

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	amount int
	coins  []int
	ret    int
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			5,
			[]int{1, 2, 5},
			4,
		},
		{
			3,
			[]int{2},
			0,
		},
		{
			500,
			[]int{3, 5, 7, 8, 9, 10, 11},
			35502874,
		},
		{
			0,
			[]int{7},
			1,
		},
	}
	for _, s := range suites {
		ret := change(s.amount, s.coins)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
