package moves

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestMoves(t *testing.T) {
	suites := []suite{
		/*{
			[]int{1, 0, 0, 1, 0, 1},
			2,
			1,
		},
		{
			[]int{1, 0, 0, 0, 0, 0, 1, 1},
			3,
			5,
		},
		{
			[]int{1, 1, 0, 1},
			2,
			0,
		},*/
		{
			[]int{0, 1, 1, 0, 0, 1, 0, 0, 0},
			3,
			2,
		},
	}

	for _, s := range suites {
		ret := minMoves(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
