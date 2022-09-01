package prices

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	prices []int
	ret    []int
}

func TestPrices(t *testing.T) {
	suites := []suite{
		{
			[]int{8, 4, 6, 2, 3},
			[]int{4, 2, 4, 2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{10, 1, 1, 6},
			[]int{9, 0, 1, 6},
		},
	}

	for _, s := range suites {
		ret := finalPrices(s.prices)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
