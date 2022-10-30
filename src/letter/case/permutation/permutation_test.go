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

func TestPermutation(t *testing.T) {
	suites := []suite{
		{
			"a1b2",
			[]string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			"3z4",
			[]string{"3z4", "3Z4"},
		},
	}
	for _, s := range suites {
		ret := letterCasePermutation(s.s)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
