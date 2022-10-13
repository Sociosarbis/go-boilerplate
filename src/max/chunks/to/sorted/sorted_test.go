package sorted

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret int
}

func TestSorted(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 3, 2, 1, 0},
			1,
		},
		{
			[]int{1, 0, 2, 3, 4},
			4,
		},
	}
	for _, s := range suites {
		ret := maxChunksToSorted(s.arr)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
