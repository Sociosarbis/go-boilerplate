package infected

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	isInfected [][]int
	ret        int
}

func TestInfected(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			4,
		},
		{
			[][]int{{0, 1, 0, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 0}},
			10,
		},
		{
			[][]int{{1, 1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 0, 0, 0, 0}},
			13,
		},
	}

	for _, s := range suites {
		ret := containVirus(s.isInfected)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
