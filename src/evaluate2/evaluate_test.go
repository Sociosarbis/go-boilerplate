package evaluate2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s         string
	knowledge [][]string
	ret       string
}

func TestEvaluate(t *testing.T) {
	suites := []suite{
		{
			"(name)is(age)yearsold",
			[][]string{{"name", "bob"}, {"age", "two"}},
			"bobistwoyearsold",
		},
		{
			"hi(name)",
			[][]string{{"a", "b"}},
			"hi?",
		},
		{
			"(a)(a)(a)aaa",
			[][]string{{"a", "yes"}},
			"yesyesyesaaa",
		},
	}

	for _, s := range suites {
		ret := evaluate(s.s, s.knowledge)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
