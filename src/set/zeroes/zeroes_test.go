package zeroes

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	matrix [][]int
	ret    [][]int
}

func TestZeroes(t *testing.T) {
	suites := []suite{}
	for _, s := range suites {
		setZeroes(s.matrix)
		if !reflect.DeepEqual(s.ret, s.matrix) {
			t.Errorf(constant.ErrTemplateStr, s.ret, s.matrix)
		}
	}
}
