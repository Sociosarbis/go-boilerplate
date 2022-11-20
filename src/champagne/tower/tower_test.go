package tower

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	poured      int
	query_row   int
	query_glass int
	ret         float64
}

func TestTower(t *testing.T) {
	suites := []suite{
		{
			1,
			1,
			1,
			0.00000,
		},
		{
			2,
			1,
			1,
			0.50000,
		},
		{
			100000009,
			33,
			17,
			1.00000,
		},
	}

	for _, s := range suites {
		ret := champagneTower(s.poured, s.query_row, s.query_glass)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
