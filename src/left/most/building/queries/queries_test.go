package queries

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	heights []int
	queries [][]int
	ret     []int
}

func TestQueries(t *testing.T) {
	suites := []suite{
		{
			[]int{6, 4, 8, 5, 2, 7},
			[][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}},
			[]int{2, 5, -1, 5, 2},
		},
		{
			[]int{5, 3, 8, 2, 6, 1, 4, 6},
			[][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}},
			[]int{7, 6, -1, 4, 6},
		},
	}

	for _, s := range suites {
		ret := leftmostBuildingQueries(s.heights, s.queries)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
