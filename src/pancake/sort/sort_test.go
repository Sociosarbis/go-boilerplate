package sort

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret []int
}

func TestSort(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 2, 4, 1},
			[]int{3, 4, 2, 3, 1, 2},
		},
		{
			[]int{1, 2, 3},
			[]int{},
		},
	}

	for _, s := range suites {
		ret := pancakeSort(s.arr)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, ret, s.ret)
		}
	}
}
