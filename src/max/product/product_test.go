package product

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestProduct(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 4, 5, 2},
			12,
		},
		{
			[]int{1, 5, 4, 5},
			16,
		},
		{
			[]int{3, 7},
			12,
		},
	}
	for _, s := range suites {
		ret := maxProduct(s.nums)
		if s.ret != ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
