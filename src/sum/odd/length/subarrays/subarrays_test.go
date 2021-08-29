package subarrays

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret int
}

func TestSumOddLengthSubarrays(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 4, 2, 5, 3},
			58,
		},
		{
			[]int{1, 2},
			3,
		},
		{
			[]int{10, 11, 12},
			66,
		},
	}

	for _, s := range suites {
		ret := sumOddLengthSubarrays(s.arr)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
