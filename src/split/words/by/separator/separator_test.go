package separator

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	words     []string
	separator byte
	ret       []string
}

func TestSeparator(t *testing.T) {
	suites := []suite{
		{
			[]string{"one.two.three", "four.five", "six"},
			'.',
			[]string{"one", "two", "three", "four", "five", "six"},
		},
		{
			[]string{"$easy$", "$problem$"},
			'$',
			[]string{"easy", "problem"},
		},
	}

	for _, s := range suites {
		ret := splitWordsBySeparator(s.words, s.separator)

		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
