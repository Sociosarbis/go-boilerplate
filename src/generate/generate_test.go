package generate

import (
	"reflect"
	"testing"
)

type suite struct {
	n   int
	ret [][]int
}

func TestGenerate(t *testing.T) {
	cases := []suite{
		{
			5,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, c := range cases {
		ret := generate(c.n)
		if !reflect.DeepEqual(ret, c.ret) {
			t.Errorf("Failed, expected %v but get %v", c.ret, ret)
		}
	}
}
