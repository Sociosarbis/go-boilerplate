package query

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums1   []int
	nums2   []int
	queries [][]int
	ret     []int64
}

func TestQuery(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 0, 1},
			[]int{0, 0, 0},
			[][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}},
			[]int64{3},
		},
		{
			[]int{1},
			[]int{5},
			[][]int{{2, 0, 0}, {3, 0, 0}},
			[]int64{5},
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0},
			[]int{48, 2, 32, 25, 30, 37, 32, 18, 48, 39, 34, 19, 46, 43, 30, 22, 20, 35, 28, 3, 5, 45, 39, 21, 46, 45, 12, 15},
			[][]int{{3, 0, 0}, {2, 3, 0}, {1, 10, 26}, {2, 4, 0}, {2, 18, 0}},
			[]int64{819},
		},
	}

	for _, s := range suites {
		ret := handleQuery(s.nums1, s.nums2, s.queries)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
