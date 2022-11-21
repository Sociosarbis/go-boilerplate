package servings

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret float64
}

func TestServings(t *testing.T) {
	suites := []suite{
		{
			50,
			0.625,
		},
		{
			100,
			0.71875,
		},
	}
	for _, s := range suites {
		ret := soupServings(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
