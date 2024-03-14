package array

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int64
}

func TestArray(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 3, 7, 9, 3},
			21,
		},
		{
			[]int{5, 3, 3},
			11,
		},
	}

	for _, s := range suites {
		ret := maxArrayValue(s.nums)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
