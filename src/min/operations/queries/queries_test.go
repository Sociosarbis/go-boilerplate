package queries

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n       int
	edges   [][]int
	queries [][]int
	ret     []int
}

func TestQueries(t *testing.T) {
	suites := []suite{
		{
			7,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 4, 2}, {4, 5, 2}, {5, 6, 2}},
			[][]int{{0, 3}, {3, 6}, {2, 6}, {0, 6}},
			[]int{0, 0, 1, 3},
		},
		{
			8,
			[][]int{{1, 2, 6}, {1, 3, 4}, {2, 4, 6}, {2, 5, 3}, {3, 6, 6}, {3, 0, 8}, {7, 0, 2}},
			[][]int{{4, 6}, {0, 4}, {6, 5}, {7, 4}},
			[]int{1, 2, 2, 3},
		},
		{
			3,
			[][]int{{2, 1, 1}, {2, 0, 2}},
			[][]int{{0, 1}, {0, 2}, {1, 2}, {0, 0}, {1, 1}, {2, 2}},
			[]int{1, 0, 0, 0, 0, 0},
		},
	}

	for _, s := range suites {
		ret := minOperationsQueries(s.n, s.edges, s.queries)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
