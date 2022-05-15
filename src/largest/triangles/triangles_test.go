package triangles

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	points [][]int
	ret    float64
}

func TestTriangles(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}},
			2.000000,
		},
		{
			[][]int{{1, 0}, {0, 0}, {0, 1}},
			0.50000,
		},
	}

	for _, s := range suites {
		ret := largestTriangleArea(s.points)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, ret, s.ret)
		}
	}
}
