package indices

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums            []int
	indexDifference int
	valueDifference int
	ret             []int
}

func TestIndices(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 1, 4, 1},
			2,
			4,
			[]int{0, 3},
		},
		{
			[]int{2, 1},
			0,
			0,
			[]int{0, 0},
		},
		{
			[]int{1, 2, 3},
			2,
			4,
			[]int{-1, -1},
		},
	}

	for _, s := range suites {
		ret := findIndices(s.nums, s.indexDifference, s.valueDifference)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
