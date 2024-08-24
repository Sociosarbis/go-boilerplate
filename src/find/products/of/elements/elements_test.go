package elements

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	queries [][]int64
	ret     []int
}

func TestElements(t *testing.T) {
	suites := []suite{
		{
			[][]int64{{0, 100000000000000, 1301}},
			[]int{430},
		},
	}

	for _, s := range suites {
		ret := findProductsOfElements(s.queries)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
