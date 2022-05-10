package win

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid      []string
	catJump   int
	mouseJump int
	ret       bool
}

func TestWin(t *testing.T) {
	suites := []suite{
		{
			[]string{"####F", "#C...", "M...."},
			1,
			2,
			true,
		},
		{
			[]string{"M.C...F"},
			1,
			4,
			true,
		},
		{
			[]string{"M.C...F"},
			1,
			3,
			false,
		},
		{
			[]string{".M...", "..#..", "#..#.", "C#.#.", "...#F"},
			3,
			1,
			true,
		},
	}

	for _, s := range suites {
		ret := canMouseWin(s.grid, s.catJump, s.mouseJump)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
