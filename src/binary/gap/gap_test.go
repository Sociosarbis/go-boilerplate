package gap

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestGap(t *testing.T) {
	suites := []suite{
		{
			22,
			2,
		},
		{
			8,
			0,
		},
		{
			5,
			2,
		},
	}

	for _, s := range suites {
		ret := binaryGap(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
