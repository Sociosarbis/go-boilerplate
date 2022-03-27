package rolls

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	rolls []int
	mean  int
	n     int
	ret   []int
}

func TestRolls(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 2, 4, 3},
			4,
			2,
			[]int{6, 6},
		},
		{
			[]int{1, 5, 6},
			3,
			4,
			[]int{3, 2, 2, 2},
		},
		{
			[]int{1, 2, 3, 4},
			6,
			4,
			[]int{},
		},
	}

	for _, s := range suites {
		ret := missingRolls(s.rolls, s.mean, s.n)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
