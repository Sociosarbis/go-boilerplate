package char

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s          string
	dictionary []string
	ret        int
}

func TestChar(t *testing.T) {
	suites := []suite{
		{
			"leetscode",
			[]string{"leet", "code", "leetcode"},
			1,
		},
		{
			"sayhelloworld",
			[]string{"hello", "world"},
			3,
		},
	}

	for _, s := range suites {
		ret := minExtraChar(s.s, s.dictionary)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
