package parentheses

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestParentheses(t *testing.T) {
	suites := []suite{
		{
			"()",
			1,
		},
		{
			"(())",
			2,
		},
		{
			"()()",
			2,
		},
	}
	for _, s := range suites {
		ret := scoreOfParentheses(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
