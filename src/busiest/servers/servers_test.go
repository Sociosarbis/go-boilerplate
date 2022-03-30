package servers

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	k       int
	arrival []int
	load    []int
	ret     []int
}

func TestServers(t *testing.T) {
	suites := []suite{
		{
			3,
			[]int{1, 2, 3, 4, 5},
			[]int{5, 2, 3, 3, 3},
			[]int{1},
		},
		{
			3,
			[]int{1, 2, 3, 4},
			[]int{1, 2, 1, 2},
			[]int{0},
		},
		{
			3,
			[]int{1, 2, 3},
			[]int{10, 12, 11},
			[]int{0, 1, 2},
		},
	}

	for _, s := range suites {
		ret := busiestServers(s.k, s.arrival, s.load)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
