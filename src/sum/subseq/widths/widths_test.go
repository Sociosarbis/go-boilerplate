package widths

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestWidths(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 1, 3},
			6,
		},
		{
			[]int{2},
			0,
		},
	}

	for _, s := range suites {
		ret := sumSubseqWidths(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
