package decrypt

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	code []int
	k    int
	ret  []int
}

func TestDecrypt(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 7, 1, 4},
			3,
			[]int{12, 10, 16, 13},
		},
		{
			[]int{1, 2, 3, 4},
			0,
			[]int{0, 0, 0, 0},
		},
		{
			[]int{2, 4, 9, 3},
			-2,
			[]int{12, 5, 6, 13},
		},
	}

	for _, s := range suites {
		ret := decrypt(s.code, s.k)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
