package two

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  []int
}

func TestTwo(t *testing.T) {
	suites := []suite{
		/*{
			[]int{1},
			[]int{2, 3},
		},
		{
			[]int{2, 3},
			[]int{1, 4},
		},*/
		{
			[]int{1, 2, 3, 4, 6, 7, 9, 10},
			[]int{5, 8},
		},
	}

	for _, s := range suites {
		ret := missingTwo(s.nums)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
