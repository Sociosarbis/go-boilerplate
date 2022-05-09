package match

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret []int
}

func TestMatch(t *testing.T) {
	suites := []suite{
		{
			"IDID",
			[]int{0, 4, 1, 3, 2},
		},
		{
			"III",
			[]int{0, 1, 2, 3},
		},
		{
			"DDI",
			[]int{3, 2, 0, 1},
		},
	}

	for _, s := range suites {
		ret := diStringMatch(s.s)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
