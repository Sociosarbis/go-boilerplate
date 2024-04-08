package operations

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestOperations(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 2, 5, 3},
			0,
		},
		{
			[]int{1, 2, 3, 5, 6},
			1,
		},
		{
			[]int{1, 10, 100, 1000},
			3,
		},
	}

	for _, s := range suites {
		ret := minOperations(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
