package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	p   string
	ret int
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			"a",
			1,
		},
		{
			"cac",
			2,
		},
		{
			"zab",
			6,
		},
	}

	for _, s := range suites {
		ret := findSubstringInWraproundString(s.p)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
