package people

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	groupSizes []int
	ret        [][]int
}

func TestPeople(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 3, 3, 3, 3, 1, 3},
			[][]int{{0, 1, 2}, {3, 4, 6}, {5}},
		},
		{
			[]int{2, 1, 3, 3, 3, 2},
			[][]int{{0, 5}, {1}, {2, 3, 4}},
		},
	}

	for _, s := range suites {
		ret := groupThePeople(s.groupSizes)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
