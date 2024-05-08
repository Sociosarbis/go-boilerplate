package plants

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	plants   []int
	capacity int
	ret      int
}

func TestPlants(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 2, 3, 3},
			5,
			14,
		},
		{
			[]int{1, 1, 1, 4, 2, 3},
			4,
			30,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			8,
			49,
		},
	}

	for _, s := range suites {
		ret := wateringPlants(s.plants, s.capacity)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
