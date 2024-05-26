package value

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	matrix [][]int
	k      int
	ret    int
}

func TestValue(t *testing.T) {
	suites := []suite{
		{
			[][]int{{5, 2}, {1, 6}},
			1,
			7,
		},
		{
			[][]int{{5, 2}, {1, 6}},
			2,
			5,
		},
		{
			[][]int{{5, 2}, {1, 6}},
			3,
			4,
		},
	}

	for _, s := range suites {
		ret := kthLargestValue(s.matrix, s.k)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
