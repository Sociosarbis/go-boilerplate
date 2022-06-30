package arrangements

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestArrangements(t *testing.T) {
	suites := []suite{
		{
			5,
			12,
		},
		{
			100,
			682289015,
		},
	}

	for _, s := range suites {
		ret := numPrimeArrangements(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
