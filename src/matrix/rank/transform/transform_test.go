package transform

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	matrix [][]int
	ret    [][]int
}

func TestTransform(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2}, {3, 4}},
			[][]int{{1, 2}, {2, 3}},
		},
		{
			[][]int{{7, 7}, {7, 7}},
			[][]int{{1, 1}, {1, 1}},
		},
		{
			[][]int{{20, -21, 14}, {-19, 4, 19}, {22, -47, 24}, {-19, 4, 19}},
			[][]int{{4, 2, 3}, {1, 3, 4}, {5, 1, 6}, {1, 3, 4}},
		},
	}

	for _, s := range suites {
		ret := matrixRankTransform(s.matrix)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
