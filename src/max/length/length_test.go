package length

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []string
	ret int
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			[]string{"un", "iq", "ue"},
			4,
		},
		{
			[]string{"cha", "r", "act", "ers"},
			6,
		},
		{
			[]string{"abcdefghijklmnopqrstuvwxyz"},
			26,
		},
	}

	for _, s := range suites {
		ret := maxLength(s.arr)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
