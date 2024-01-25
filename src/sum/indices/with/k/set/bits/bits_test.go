package bits

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestBits(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 10, 1, 5, 2},
			1,
			13,
		},
		{
			[]int{4, 3, 2, 1},
			2,
			1,
		},
	}

	for _, s := range suites {
		ret := sumIndicesWithKSetBits(s.nums, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
