package sum

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int64
}

func TestSum(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 4, -2},
			5,
			2,
		},
		{
			[]int{1, -2, 3, 4, -10, 12},
			16,
			10,
		},
	}

	for _, s := range suites {
		ret := kSum(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
