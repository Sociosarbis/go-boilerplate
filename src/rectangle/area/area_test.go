package area

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	rectangles [][]int
	ret        int
}

func TestArea(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}},
			6,
		},
		{
			[][]int{{0, 0, 1000000000, 1000000000}},
			49,
		},
	}

	for _, s := range suites {
		ret := rectangleArea(s.rectangles)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
