package averages

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  float64
}

func TestAverages(t *testing.T) {
	suites := []suite{
		{
			[]int{9, 1, 2, 3, 9},
			3,
			20.00000,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			4,
			20.50000,
		},
	}

	for _, s := range suites {
		ret := largestSumOfAverages(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
