package competitive

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  []int
}

func TestCompetitive(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 5, 2, 6},
			2,
			[]int{2, 6},
		},
		{
			[]int{2, 4, 3, 3, 5, 4, 9, 6},
			4,
			[]int{2, 3, 3, 4},
		},
	}

	for _, s := range suites {
		ret := mostCompetitive(s.nums, s.k)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
