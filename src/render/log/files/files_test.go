package files

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	logs []string
	ret  []string
}

func TestFiles(t *testing.T) {
	suites := []suite{
		{
			[]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			[]string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			[]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			[]string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
	}

	for _, s := range suites {
		ret := reorderLogFiles(s.logs)

		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
