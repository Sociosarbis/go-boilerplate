package subsequence

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  []int
}

func TestSubsequence(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 3, 10, 9, 8},
			[]int{10, 9},
		},
		{
			[]int{4, 4, 7, 6, 7},
			[]int{7, 7, 6},
		},
	}

	for _, s := range suites {
		ret := minSubsequence(s.nums)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
