package bricks

import (
	"reflect"
	"testing"
)

type suite struct {
	grid [][]int
	hits [][]int
	ret  []int
}

func TestBricks(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 0, 0, 0}, {1, 1, 1, 0}},
			[][]int{{1, 0}},
			[]int{2},
		},
		{
			[][]int{{1, 0, 0, 0}, {1, 1, 0, 0}},
			[][]int{{1, 1}, {1, 0}},
			[]int{0, 0},
		},
		{
			[][]int{{1, 0, 1}, {1, 1, 1}},
			[][]int{{0, 0}, {0, 2}, {1, 1}},
			[]int{0, 3, 0},
		},
		{
			[][]int{{0, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 0}, {0, 0, 1, 1, 0}, {0, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
			[][]int{{6, 0}, {1, 0}, {4, 3}, {1, 2}, {7, 1}, {6, 3}, {5, 2}, {5, 1}, {2, 4}, {4, 4}, {7, 3}},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, s := range suites {
		ret := hitBricks(s.grid, s.hits)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
