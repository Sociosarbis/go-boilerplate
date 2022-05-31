package order

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	words []string
	ret   string
}

func TestOrder(t *testing.T) {
	suites := []suite{
		{
			[]string{"wrt", "wrf", "er", "ett", "rftt"},
			"wertf",
		},
		{
			[]string{"z", "x"},
			"zx",
		},
		{
			[]string{"z", "x", "z"},
			"",
		},
		{
			[]string{"ac", "ab", "zc", "zb"},
			"acbz",
		},
	}

	for _, s := range suites {
		ret := alienOrder(s.words)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, ret, s.ret)
		}
	}
}
