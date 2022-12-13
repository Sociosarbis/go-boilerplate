package sum

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestSum(t *testing.T) {
	suites := []suite{
		{
			"aabcb",
			5,
		},
		{
			"aabcbaa",
			17,
		},
	}

	for _, s := range suites {
		ret := beautySum(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
