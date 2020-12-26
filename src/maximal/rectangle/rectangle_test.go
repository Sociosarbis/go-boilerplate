package rectangle

import (
	"testing"
)

type suite struct {
	matrix [][]byte
	ret    int
}

func TestRectangle(t *testing.T) {
	suites := []suite{
		{
			[][]byte{
				{49, 48, 49, 48, 48},
				{49, 48, 49, 49, 49},
				{49, 49, 49, 49, 49},
				{49, 48, 48, 49, 48},
			},
			6,
		},
		{
			[][]byte{
				{48},
			},
			0,
		},
		{
			[][]byte{
				{49},
			},
			1,
		},
		{
			[][]byte{
				{48, 48},
			},
			0,
		},
	}

	for _, s := range suites {
		ret := maximalRectangle(s.matrix)
		if maximalRectangle(s.matrix) != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
