package order

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret []int
}

func TestOrder(t *testing.T) {
	suites := []suite{
		{
			13,
			[]int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			2,
			[]int{1, 2},
		},
	}

	for _, s := range suites {
		ret := lexicalOrder(s.n)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
