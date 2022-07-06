package evaluate

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	expression string
	ret        int
}

func TestEvaluate(t *testing.T) {
	suites := []suite{
		{
			"(let x 2 (mult x (let x 3 y 4 (add x y))))",
			14,
		},
		{
			"(let x 3 x 2 x)",
			2,
		},
		{
			"(let x 1 y 2 x (add x y) (add x y))",
			5,
		},
		{
			"(let x 7 -12)",
			-12,
		},
	}

	for _, s := range suites {
		ret := evaluate(s.expression)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
