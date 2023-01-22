package groups

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	batchSize int
	groups    []int
	ret       int
}

func TestGroups(t *testing.T) {
	suites := []suite{
		{
			3,
			[]int{1, 2, 3, 4, 5, 6},
			4,
		},
		{
			4,
			[]int{1, 3, 2, 5, 2, 2, 1, 6},
			4,
		},
		{
			9,
			[]int{298148512, 266520299, 406786559, 804851972, 600623393, 4304578, 203837753, 567325621, 471128999, 621785239, 22585811, 33080261, 673801543, 650287622, 260102349, 896931288, 279598926, 221841310, 185598694, 305781935, 159354503, 582332530, 215228600, 293373862},
			12,
		},
	}

	for _, s := range suites {
		ret := maxHappyGroups(s.batchSize, s.groups)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
