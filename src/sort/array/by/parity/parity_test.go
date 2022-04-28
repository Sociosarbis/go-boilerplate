package parity

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  []int
}

func TestParity(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 1, 2, 4},
			[]int{4, 2, 1, 3},
		},
		{
			[]int{0},
			[]int{0},
		},
	}

	for _, s := range suites {
		ret := sortArrayByParity(s.nums)

		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
