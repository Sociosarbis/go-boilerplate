package words

import (
	"reflect"
	"testing"
)

type suite struct {
	words   []string
	puzzles []string
	ret     []int
}

func TestWords(t *testing.T) {
	suites := []suite{
		{
			[]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
			[]string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			[]int{1, 1, 3, 2, 4, 0},
		},
	}

	for _, s := range suites {
		ret := findNumOfValidWords(s.words, s.puzzles)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
