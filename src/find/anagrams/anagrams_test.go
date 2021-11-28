package anagrams

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	p   string
	ret []int
}

func TestAnagrams(t *testing.T) {
	suites := []suite{
		{
			"cbaebabacd",
			"abc",
			[]int{0, 6},
		},
		{
			"abab",
			"ab",
			[]int{0, 1, 2},
		},
	}

	for _, s := range suites {
		ret := findAnagrams(s.s, s.p)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
