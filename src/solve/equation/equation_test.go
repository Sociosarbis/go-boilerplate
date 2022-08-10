package equation

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	equation string
	ret      string
}

func TestEquation(t *testing.T) {
	suites := []suite{
		{
			"x+5-3+x=6+x-2",
			"x=2",
		},
		{
			"x=x",
			"Infinite solutions",
		},
		{
			"2x=x",
			"x=0",
		},
	}

	for _, s := range suites {
		ret := solveEquation(s.equation)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
