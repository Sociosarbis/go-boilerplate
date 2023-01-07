package operations4

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	x    int
	ret  int
}

func TestOperations(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 1, 4, 2, 3},
			5,
			2,
		},
		{
			[]int{5, 6, 7, 8, 9},
			4,
			-1,
		},
		{
			[]int{3, 2, 20, 1, 1, 3},
			10,
			5,
		},
		{
			[]int{1, 1},
			3,
			-1,
		},
	}

	for _, s := range suites {
		ret := minOperations(s.nums, s.x)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
