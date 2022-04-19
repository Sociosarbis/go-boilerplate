package char

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	c   byte
	ret []int
}

func TestChar(t *testing.T) {
	suites := []suite{
		{
			"loveleetcode",
			'e',
			[]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			"aaab",
			'b',
			[]int{3, 2, 1, 0},
		},
	}

	for _, s := range suites {
		ret := shortestToChar(s.s, s.c)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
