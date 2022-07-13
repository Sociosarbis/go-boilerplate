package asteroid

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	asteroids []int
	ret       []int
}

func TestCollision(t *testing.T) {
	suites := []suite{
		{
			[]int{5, 10, -5},
			[]int{5, 10},
		},
		{
			[]int{8, -8},
			[]int{},
		},
		{
			[]int{10, 2, -5},
			[]int{10},
		},
	}

	for _, s := range suites {
		ret := asteroidCollision(s.asteroids)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
