package expr

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	expression string
	ret        bool
}

func TestExpr(t *testing.T) {
	suites := []suite{
		{
			"&(|(f))",
			false,
		},
		{
			"|(f,f,f,t)",
			true,
		},
		{
			"!(&(f,t))",
			true,
		},
	}

	for _, s := range suites {
		ret := parseBoolExpr(s.expression)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
