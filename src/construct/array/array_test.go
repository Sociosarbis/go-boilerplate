package array

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	k   int
	ret []int
}

func TestArray(t *testing.T) {
	suites := []suite{
		{
			3,
			1,
			[]int{1, 2, 3},
		},
		{
			3,
			2,
			[]int{1, 3, 2},
		},
	}

	for _, s := range suites {
		ret := constructArray(s.n, s.k)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
