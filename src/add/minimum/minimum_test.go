package minimum

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	word string
	ret  int
}

func TestMinimum(t *testing.T) {
	suites := []suite{
		{
			"b",
			2,
		},
		{
			"aaa",
			6,
		},
		{
			"abc",
			0,
		},
	}

	for _, s := range suites {
		ret := addMinimum(s.word)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
