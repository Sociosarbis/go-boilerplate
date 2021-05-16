package xor

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestNumbers(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 10, 5, 25, 2, 8},
			28,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{2, 4},
			6,
		},
		{
			[]int{8, 10, 2},
			10,
		},
	}
	for _, su := range suites {
		ret := findMaximumXOR(su.nums)
		if ret != su.ret {
			t.Errorf(constant.ErrTemplateStr, su.ret, ret)
		}
	}
}
