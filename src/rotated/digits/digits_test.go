package digits

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestDigits(t *testing.T) {
	suites := []suite{
		{
			10,
			4,
		},
		{
			1,
			0,
		},
		{
			2,
			1,
		},
	}

	for _, s := range suites {
		ret := rotatedDigits(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
