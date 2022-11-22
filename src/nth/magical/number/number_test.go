package number

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	a   int
	b   int
	ret int
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			1,
			2,
			3,
			2,
		},
		{
			4,
			2,
			3,
			6,
		},
	}

	for _, s := range suites {
		ret := nthMagicalNumber(s.n, s.a, s.b)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
