package winner

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	player1 []int
	player2 []int
	ret     int
}

func TestWinner(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 10, 7, 9},
			[]int{6, 5, 2, 3},
			1,
		},
		{
			[]int{3, 5, 7, 6},
			[]int{8, 10, 10, 2},
			2,
		},
		{
			[]int{2, 3},
			[]int{4, 1},
			0,
		},
	}

	for _, s := range suites {
		ret := isWinner(s.player1, s.player2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
