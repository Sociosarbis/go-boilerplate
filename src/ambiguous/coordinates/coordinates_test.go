package coordinates

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret []string
}

func TestCoordinates(t *testing.T) {
	suites := []suite{
		{
			"(123)",
			[]string{"(1, 2.3)", "(1, 23)", "(1.2, 3)", "(12, 3)"},
		},
		{
			"(0123)",
			[]string{"(0, 1.23)", "(0, 12.3)", "(0, 123)", "(0.1, 2.3)", "(0.1, 23)", "(0.12, 3)"},
		},
		{
			"(00011)",
			[]string{"(0, 0.011)", "(0.001, 1)"},
		},
		{
			"(0010)",
			[]string{"(0.01, 0)"},
		},
	}

	for _, s := range suites {
		ret := ambiguousCoordinates(s.s)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
