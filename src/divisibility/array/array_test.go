package array

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	word string
	m    int
	ret  []int
}

func TestArray(t *testing.T) {
	suites := []suite{
		{
			"998244353",
			3,
			[]int{1, 1, 0, 0, 0, 1, 1, 0, 0},
		},
		{
			"1010",
			10,
			[]int{0, 1, 0, 1},
		},
	}

	for _, s := range suites {
		ret := divisibilityArray(s.word, s.m)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
