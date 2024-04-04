package ancestors

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n     int
	edges [][]int
	ret   [][]int
}

func TestAncestors(t *testing.T) {
	suites := []suite{
		{
			8,
			[][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}},
			[][]int{{}, {}, {}, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}},
		},
		{
			5,
			[][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			[][]int{{}, {0}, {0, 1}, {0, 1, 2}, {0, 1, 2, 3}},
		},
	}

	for _, s := range suites {
		ret := getAncestors(s.n, s.edges)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
