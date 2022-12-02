package operation2

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	boxes string
	ret   []int
}

func TestOperations(t *testing.T) {
	suites := []suite{
		{
			"110",
			[]int{1, 1, 3},
		},
		{
			"001011",
			[]int{11, 8, 5, 4, 3, 4},
		},
	}

	for _, s := range suites {
		ret := minOperations(s.boxes)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
