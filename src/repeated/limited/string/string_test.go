package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s           string
	repeatLimit int
	ret         string
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			"cczazcc",
			3,
			"zzcccac",
		},
		{
			"aababab",
			2,
			"bbabaa",
		},
	}

	for _, s := range suites {
		ret := repeatLimitedString(s.s, s.repeatLimit)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
