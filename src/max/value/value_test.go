package value

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n      int
	index  int
	maxSum int
	ret    int
}

func TestValue(t *testing.T) {
	suites := []suite{
		{
			4,
			2,
			6,
			2,
		},
		{
			6,
			1,
			10,
			3,
		},
	}

	for _, s := range suites {
		ret := maxValue(s.n, s.index, s.maxSum)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
