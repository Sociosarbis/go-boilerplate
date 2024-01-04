package rows

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	matrix    [][]int
	numSelect int
	ret       int
}

func TestRows(t *testing.T) {
	suites := []suite{{
		[][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}},
		2,
		3,
	}, {
		[][]int{{1}, {0}},
		1,
		2,
	}}

	for _, s := range suites {
		ret := maximumRows(s.matrix, s.numSelect)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
