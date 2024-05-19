package winner

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	k   int
	ret int
}

func TestWinner(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 1, 3, 5, 4, 6, 7},
			2,
			5,
		},
		{
			[]int{3, 2, 1},
			10,
			3,
		},
	}

	for _, s := range suites {
		ret := getWinner(s.arr, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
