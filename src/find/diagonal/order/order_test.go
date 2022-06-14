package order

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	mat [][]int
	ret []int
}

func TestOrder(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 3, 4},
		},
	}

	for _, s := range suites {
		ret := findDiagonalOrder(s.mat)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
