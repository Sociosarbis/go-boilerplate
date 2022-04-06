package trees

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n     int
	edges [][]int
	ret   []int
}

func TestTrees(t *testing.T) {
	suites := []suite{
		{
			4,
			[][]int{{1, 0}, {1, 2}, {1, 3}},
			[]int{1},
		},
		{
			6,
			[][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			[]int{3, 4},
		},
	}

	for _, su := range suites {
		ret := findMinHeightTrees(su.n, su.edges)
		if !reflect.DeepEqual(su.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, su.ret, ret)
		}
	}
}
