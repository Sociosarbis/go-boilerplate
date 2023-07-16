package tree

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

func TestTree(t *testing.T) {
	suites := []suite{
		{
			6,
			[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			[]int{8, 12, 6, 10, 10, 10},
		},
		{
			1,
			[][]int{},
			[]int{0},
		},
		{
			2,
			[][]int{{1, 0}},
			[]int{1, 1},
		},
	}

	for _, s := range suites {
		ret := sumOfDistancesInTree(s.n, s.edges)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
