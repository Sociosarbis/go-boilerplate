package width

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  []int
}

func TestWidth(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1}, {22}, {333}},
			[]int{3},
		},
		{
			[][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}},
			[]int{3, 1, 2},
		},
	}

	for _, s := range suites {
		ret := findColumnWidth(s.grid)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
