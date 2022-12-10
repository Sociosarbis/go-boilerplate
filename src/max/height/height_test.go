package height

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	cuboids [][]int
	ret     int
}

func TestHeight(t *testing.T) {
	suites := []suite{
		{
			[][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}},
			190,
		},
		{
			[][]int{{38, 25, 45}, {76, 35, 3}},
			76,
		},
		{
			[][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}},
			102,
		},
	}

	for _, s := range suites {
		ret := maxHeight(s.cuboids)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
