package special

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret string
}

func TestSpecial(t *testing.T) {
	suites := []suite{
		{
			"11011000",
			"11100100",
		},
	}

	for _, s := range suites {
		ret := makeLargestSpecial(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
