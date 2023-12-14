package stamp

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid        [][]int
	stampHeight int
	stampWidth  int
	ret         bool
}

func TestStamp(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}},
			4,
			3,
			true,
		},
		{
			[][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
			2,
			2,
			false,
		},
		{
			[][]int{{0}, {0}, {0}, {0}, {1}, {1}, {0}, {0}, {1}},
			9,
			1,
			false,
		},
		{
			[][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 1}, {0, 0, 0, 1, 1}},
			2,
			2,
			false,
		},
	}

	for _, s := range suites {
		ret := possibleToStamp(s.grid, s.stampHeight, s.stampWidth)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
