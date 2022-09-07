package spaces

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	text string
	ret  string
}

func TestSpaces(t *testing.T) {
	suites := []suite{
		{
			"  this   is  a sentence ",
			"this   is   a   sentence",
		},
		{
			" practice   makes   perfect",
			"practice   makes   perfect ",
		},
	}

	for _, s := range suites {
		ret := reorderSpaces(s.text)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
