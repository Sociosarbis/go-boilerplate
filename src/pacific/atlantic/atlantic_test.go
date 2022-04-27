package atlantic

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	heights [][]int
	ret     [][]int
}

func TestAtlantic(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			[][]int{{2, 1}, {1, 2}},
			[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}

	for _, s := range suites {
		ret := pacificAtlantic(s.heights)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
