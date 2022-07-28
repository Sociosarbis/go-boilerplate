package transform

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret []int
}

func TestTransform(t *testing.T) {
	suites := []suite{
		{
			[]int{40, 10, 20, 30},
			[]int{4, 1, 2, 3},
		},
		{
			[]int{100, 100, 100},
			[]int{1, 1, 1},
		},
		{
			[]int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			[]int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
	}
	for _, s := range suites {
		ret := arrayRankTransform(s.arr)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
