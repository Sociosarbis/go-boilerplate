package lines

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	widths []int
	s      string
	ret    []int
}

func TestLines(t *testing.T) {
	suites := []suite{
		{
			[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"abcdefghijklmnopqrstuvwxyz",
			[]int{3, 60},
		},
		{
			[]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"bbbcccdddaaa",
			[]int{2, 4},
		},
	}

	for _, s := range suites {
		ret := numberOfLines(s.widths, s.s)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
