package sum

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestSum(t *testing.T) {
	suites := []suite{
		{
			[]int{10, 20, 30, 5, 10, 50},
			65,
		},
		{
			[]int{10, 20, 30, 40, 50},
			150,
		},
		{
			[]int{12, 17, 15, 13, 10, 11, 12},
			33,
		},
	}

	for _, s := range suites {
		ret := maxAscendingSum(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
