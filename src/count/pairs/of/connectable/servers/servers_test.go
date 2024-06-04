package servers

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	edges       [][]int
	signalSpeed int
	ret         []int
}

func TestServers(t *testing.T) {
	suites := []suite{
		{
			edges:       [][]int{{0, 1, 1}, {1, 2, 5}, {2, 3, 13}, {3, 4, 9}, {4, 5, 2}},
			signalSpeed: 1,
			ret:         []int{0, 4, 6, 6, 4, 0},
		},
		{
			edges:       [][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}},
			signalSpeed: 3,
			ret:         []int{2, 0, 0, 0, 0, 0, 2},
		},
	}

	for _, s := range suites {
		ret := countPairsOfConnectableServers(s.edges, s.signalSpeed)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
