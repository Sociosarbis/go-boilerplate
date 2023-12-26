package students

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	seats [][]string
	ret   int
}

func TestStudents(t *testing.T) {
	suites := []suite{
		{
			[][]string{{"#", ".", "#", "#", ".", "#"},
				{".", "#", "#", "#", "#", "."},
				{"#", ".", "#", "#", ".", "#"}},
			4,
		},
		{
			[][]string{{".", "#"},
				{"#", "#"},
				{"#", "."},
				{"#", "#"},
				{".", "#"}},
			3,
		},
		{
			[][]string{{"#", ".", ".", ".", "#"},
				{".", "#", ".", "#", "."},
				{".", ".", "#", ".", "."},
				{".", "#", ".", "#", "."},
				{"#", ".", ".", ".", "#"}},
			10,
		},
		{
			[][]string{{"#", "#", "#"}, {".", "#", "#"}},
			1,
		},
		{
			[][]string{{"#", "#", "#", ".", "#"}, {".", ".", "#", ".", "."}, {"#", ".", "#", ".", "#"}, {".", ".", ".", ".", "."}, {".", ".", ".", "#", "."}},
			9,
		},
		{
			[][]string{{"#", "#", ".", ".", "."}, {"#", ".", ".", "#", "."}, {".", "#", ".", ".", "#"}, {".", "#", "#", "#", "#"}, {".", ".", "#", ".", "."}},
			9,
		},
	}

	for _, s := range suites {
		seats := make([][]byte, len(s.seats))
		for i := range seats {
			seats[i] = make([]byte, len(s.seats[i]))
			for j := range seats[i] {
				seats[i][j] = s.seats[i][j][0]
			}
		}
		ret := maxStudents(seats)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
