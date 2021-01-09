package profit

import (
	"testing"
)

type suite struct {
	prices []int
	ret    int
}

func TestMaxProfit(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 3, 5, 0, 0, 3, 1, 4},
			6,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{1},
			0,
		},
	}

	for _, s := range suites {
		ret := maxProfit(s.prices)
		if ret != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
