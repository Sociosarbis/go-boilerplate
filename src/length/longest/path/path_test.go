package path

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	input string
	ret   int
}

func TestPath(t *testing.T) {
	suites := []suite{
		{
			"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			20,
		},
		{
			"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			32,
		},
		{
			"a",
			0,
		},
	}

	for _, s := range suites {
		ret := lengthLongestPath(s.input)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
