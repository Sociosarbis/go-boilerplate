package collection

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	garbage []string
	travel  []int
	ret     int
}

func TestCollection(t *testing.T) {
	suites := []suite{
		{
			[]string{"G", "P", "GP", "GG"},
			[]int{2, 4, 3},
			21,
		},
		{
			[]string{"MMM", "PGM", "GP"},
			[]int{3, 10},
			37,
		},
	}

	for _, s := range suites {
		ret := garbageCollection(s.garbage, s.travel)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
