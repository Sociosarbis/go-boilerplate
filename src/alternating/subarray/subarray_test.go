package subarray

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestSubarray(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 3, 4, 3, 4},
			4,
		},
		{
			[]int{4, 5, 6},
			2,
		},
	}

	for _, s := range suites {
		ret := alternatingSubarray(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
