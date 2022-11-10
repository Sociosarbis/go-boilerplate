package keys

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid []string
	ret  int
}

func TestKey(t *testing.T) {
	suites := []suite{
		{
			[]string{"@.a..", "###.#", "b.A.B"},
			8,
		},
		{
			[]string{"@.a..", "###.#", "b.A.B"},
			8,
		},
		{
			[]string{"@Aa"},
			-1,
		},
		{
			[]string{"bAa@B"},
			3,
		},
		{
			[]string{"@...a", ".###A", "b.BCc"},
			10,
		},
		{
			[]string{".#.b.", "A.#aB", "#d...", "@.cC.", "D...#"},
			8,
		},
	}

	for _, s := range suites {
		ret := shortestPathAllKeys(s.grid)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
