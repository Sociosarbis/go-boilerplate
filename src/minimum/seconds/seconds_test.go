package seconds

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestSeconds(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 1, 2},
			1,
		},
		{
			[]int{2, 1, 3, 3, 2},
			2,
		},
		{
			[]int{5, 5, 5, 5},
			0,
		},
	}

	for _, s := range suites {
		ret := minimumSeconds(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
