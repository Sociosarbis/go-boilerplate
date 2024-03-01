package partition

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  bool
}

func TestPartition(t *testing.T) {
	suites := []suite{
		// {
		// 	[]int{4, 4, 4, 5, 6},
		// 	true,
		// },
		{
			[]int{1, 1, 1, 2},
			false,
		},
	}

	for _, s := range suites {
		ret := validPartition(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
