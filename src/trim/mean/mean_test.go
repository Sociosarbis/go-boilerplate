package mean

import (
	"math"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret float64
}

func TestMean(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3},
			2.00000,
		},
		{
			[]int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0},
			4.00000,
		},
		{
			[]int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4},
			4.77778,
		},
	}
	for _, s := range suites {
		ret := trimMean(s.arr)
		if math.Abs(ret-s.ret) > 0.00001 {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
