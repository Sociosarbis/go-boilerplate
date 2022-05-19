package move2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestMove(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3},
			2,
		},
		{
			[]int{1, 10, 2, 9},
			16,
		},
	}

	for _, s := range suites {
		ret := minMoves2(s.nums)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
