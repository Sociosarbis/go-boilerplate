package burgers

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	tomatoSlices int
	cheeseSlices int
	ret          []int
}

func TestBurgers(t *testing.T) {
	suites := []suite{
		{
			16,
			7,
			[]int{1, 6},
		},
		{
			17,
			4,
			nil,
		},
		{
			4,
			17,
			nil,
		},
	}

	for _, s := range suites {
		ret := numOfBurgers(s.tomatoSlices, s.cheeseSlices)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
