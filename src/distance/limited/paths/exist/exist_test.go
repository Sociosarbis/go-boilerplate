package exist

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n        int
	edgeList [][]int
	queries  [][]int
	ret      []bool
}

func TestExist(t *testing.T) {
	suites := []suite{
		{
			3,
			[][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}},
			[][]int{{0, 1, 2}, {0, 2, 5}},
			[]bool{false, true},
		},
		{
			5,
			[][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}},
			[][]int{{0, 4, 14}, {1, 4, 13}},
			[]bool{true, false},
		},
	}

	for _, s := range suites {
		ret := distanceLimitedPathsExist(s.n, s.edgeList, s.queries)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
