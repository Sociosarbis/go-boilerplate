package chessboard

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	board [][]int
	ret   int
}

func TestChessboard(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}},
			2,
		},
		{
			[][]int{{0, 1}, {1, 0}},
			0,
		},
		{
			[][]int{{1, 0}, {1, 0}},
			-1,
		},
	}

	for _, s := range suites {
		ret := movesToChessboard(s.board)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
