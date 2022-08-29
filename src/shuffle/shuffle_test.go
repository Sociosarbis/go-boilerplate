package shuffle

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	n    int
	ret  []int
}

func TestShuffle(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 5, 1, 3, 4, 7},
			3,
			[]int{2, 3, 5, 4, 1, 7},
		},
		{
			[]int{1, 2, 3, 4, 4, 3, 2, 1},
			4,
			[]int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			[]int{1, 1, 2, 2},
			2,
			[]int{1, 2, 1, 2},
		},
	}

	for _, s := range suites {
		ret := shuffle(s.nums, s.n)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
