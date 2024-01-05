package count

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	heights []int
	ret     []int
}

func TestCount(t *testing.T) {
	suites := []suite{
		{
			[]int{10, 6, 8, 5, 11, 9},
			[]int{3, 1, 2, 1, 1, 0},
		},
		{
			[]int{5, 1, 2, 3, 10},
			[]int{4, 1, 1, 1, 0},
		},
	}

	for _, s := range suites {
		ret := canSeePersonsCount(s.heights)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
