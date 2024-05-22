package winners

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	matches [][]int
	ret     [][]int
}

func TestWinners(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
			[][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
		{
			[][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}},
			[][]int{{1, 2, 5, 6}, {}},
		},
	}

	for _, s := range suites {
		ret := findWinners(s.matches)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
