package length

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	strs []string
	ret  int
}

func TestLength(t *testing.T) {
	suites := []suite{
		{
			[]string{"aba", "cdc", "eae"},
			3,
		},
		{
			[]string{"aaa", "aaa", "aa"},
			-1,
		},
	}

	for _, s := range suites {
		ret := findLUSlength(s.strs)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
