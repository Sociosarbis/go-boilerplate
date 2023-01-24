package points

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	points  [][]int
	queries [][]int
	ret     []int
}

func TestPoints(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}},
			[][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}},
			[]int{3, 2, 2},
		},
		{
			[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			[][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}},
			[]int{2, 3, 2, 4},
		},
	}

	for _, s := range suites {
		ret := countPoints(s.points, s.queries)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
