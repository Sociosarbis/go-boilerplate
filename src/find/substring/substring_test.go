package substring

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s     string
	words []string
	ret   []int
}

func TestSubstring(t *testing.T) {
	suites := []suite{
		{
			"barfoothefoobarman",
			[]string{"foo", "bar"},
			[]int{0, 9},
		},
		{
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"},
			[]int{},
		},
		{
			"barfoofoobarthefoobarman",
			[]string{"bar", "foo", "the"},
			[]int{6, 9, 12},
		},
	}

	for _, s := range suites {
		ret := findSubstring(s.s, s.words)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
