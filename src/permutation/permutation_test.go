package permutation

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret []string
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			"abc",
			[]string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{"aab",
			[]string{"aab", "aba", "baa"},
		},
	}

	for _, s := range suites {
		ret := permutation(s.s)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
