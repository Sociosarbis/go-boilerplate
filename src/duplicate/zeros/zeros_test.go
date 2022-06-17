package zeros

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret []int
}

func TestZeros(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 0, 2, 3, 0, 4, 5, 0},
			[]int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
	}

	for _, s := range suites {
		duplicateZeros(s.arr)
		if !reflect.DeepEqual(s.arr, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, s.arr)
		}
	}
}
