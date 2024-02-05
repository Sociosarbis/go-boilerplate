package result

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestResult(t *testing.T) {
	suites := []suite{
		{
			[]int{1, -1, -2, 4, -7, 3},
			2,
			7,
		},
		{
			[]int{10, -5, -2, 4, 0, 3},
			3,
			17,
		},
		{
			[]int{1, -5, -20, 4, -1, 3, -6, -3},
			2,
			0,
		},
	}

	for _, s := range suites {
		ret := maxResult(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
