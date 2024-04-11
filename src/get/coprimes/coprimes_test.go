package coprimes

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums  []int
	edges [][]int
	ret   []int
}

func TestCoprimes(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 3, 3, 2},
			[][]int{{0, 1}, {1, 2}, {1, 3}},
			[]int{-1, 0, 0, 1},
		},
		{
			[]int{5, 6, 10, 2, 3, 6, 15},
			[][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
			[]int{-1, 0, -1, 0, 0, 0, -1},
		},
	}

	for _, s := range suites {
		ret := getCoprimes(s.nums, s.edges)

		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
