package tower

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestTower(t *testing.T) {
	suites := []suite{
		{
			[]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150},
			1,
		},
		{
			[]int{-200, -300, 400, 0},
			-1,
		},
	}

	for _, s := range suites {
		ret := magicTower(s.nums)
		if s.ret != ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
