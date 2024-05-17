package assignment

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	difficulty []int
	profit     []int
	worker     []int
	ret        int
}

func TestAssignment(t *testing.T) {
	suites := []suite{
		// {
		// 	[]int{2, 4, 6, 8, 10},
		// 	[]int{10, 20, 30, 40, 50},
		// 	[]int{4, 5, 6, 7},
		// 	100,
		// },
		// {
		// 	[]int{85, 47, 57},
		// 	[]int{24, 66, 99},
		// 	[]int{40, 25, 25},
		// 	0,
		// },
		// {
		// 	[]int{5, 50, 92, 21, 24, 70, 17, 63, 30, 53},
		// 	[]int{68, 100, 3, 99, 56, 43, 26, 93, 55, 25},
		// 	[]int{96, 3, 55, 30, 11, 58, 68, 36, 26, 1},
		// 	765,
		// },
		{
			[]int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63},
			[]int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77},
			[]int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16},
			1392,
		},
	}

	for _, s := range suites {
		ret := maxProfitAssignment(s.difficulty, s.profit, s.worker)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
