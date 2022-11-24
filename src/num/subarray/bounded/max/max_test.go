package max

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums  []int
	left  int
	right int
	ret   int
}

func TestMax(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 1, 4, 3},
			2,
			3,
			3,
		},
		{
			[]int{2, 9, 2, 5, 6},
			2,
			8,
			7,
		},
	}

	for _, s := range suites {
		ret := numSubarrayBoundedMax(s.nums, s.left, s.right)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
