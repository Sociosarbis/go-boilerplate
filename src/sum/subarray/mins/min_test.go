package mins

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret int
}

func TestMins(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 1, 2, 4},
			17,
		},
		{
			[]int{11, 81, 94, 43, 3},
			444,
		},
	}
	for _, s := range suites {
		ret := sumSubarrayMins(s.arr)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
