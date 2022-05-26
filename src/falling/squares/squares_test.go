package squares

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	positions [][]int
	ret       []int
}

func TestSquares(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2}, {2, 3}, {6, 1}},
			[]int{2, 5, 5},
		},
		{
			[][]int{{100, 100}, {200, 100}},
			[]int{100, 100},
		},
		{[][]int{{1, 37}, {59, 2}, {77, 64}, {21, 92}, {6, 44}},
			[]int{37, 37, 64, 156, 200},
		},
	}

	for _, s := range suites {
		ret := fallingSquares(s.positions)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
