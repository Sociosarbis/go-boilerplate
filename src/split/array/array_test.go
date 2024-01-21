package array

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestArray(t *testing.T) {
	suites := []suite{
		// {
		// 	[]int{7, 2, 5, 10, 8},
		// 	2,
		// 	18,
		// },
		// {
		// 	[]int{1, 2, 3, 4, 5},
		// 	2,
		// 	9,
		// },
		{
			[]int{1, 4, 4},
			3,
			4,
		},
	}

	for _, s := range suites {
		ret := splitArray(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
