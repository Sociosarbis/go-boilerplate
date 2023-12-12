package element

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  []int
}

func TestElement(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 4, 0, 9, 6},
			[]int{9, 6, 6, -1, -1},
		},
		{
			[]int{3, 3},
			[]int{-1, -1},
		},
	}

	for _, s := range suites {
		ret := secondGreaterElement(s.nums)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
