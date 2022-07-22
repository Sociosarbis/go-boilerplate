package two

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	intervals [][]int
	ret       int
}

func TestTwo(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			3,
		},
		{
			[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			5,
		},
		{
			[][]int{{2, 10}, {3, 7}, {3, 15}, {4, 11}, {6, 12}, {6, 16}, {7, 8}, {7, 11}, {7, 15}, {11, 12}},
			5,
		},
		{
			[][]int{{1, 3}, {3, 7}, {5, 7}, {7, 8}},
			5,
		},
		{
			[][]int{{12, 19}, {18, 25}, {4, 6}, {19, 24}, {19, 22}},
			5,
		},
		{
			[][]int{{1, 3}, {4, 9}, {0, 10}, {6, 7}, {1, 2}, {0, 6}, {7, 9}, {0, 1}, {2, 5}, {6, 8}},
			7,
		},
		{
			[][]int{{1, 3}, {7, 10}, {5, 9}, {5, 8}, {6, 7}, {0, 3}, {5, 8}, {3, 4}, {6, 10}, {4, 6}},
			6,
		},
	}

	for _, s := range suites {
		ret := intersectionSizeTwo(s.intervals)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
