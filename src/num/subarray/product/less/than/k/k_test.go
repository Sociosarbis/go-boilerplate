package k

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestK(t *testing.T) {
	suites := []suite{
		{
			[]int{10, 5, 2, 6},
			100,
			8,
		},
		{
			[]int{1, 2, 3},
			0,
			0,
		},
		{
			[]int{57, 44, 92, 28, 66, 60, 37, 33, 52, 38, 29, 76, 8, 75, 22},
			18,
			1,
		},
	}

	for _, s := range suites {
		ret := numSubarrayProductLessThanK(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
