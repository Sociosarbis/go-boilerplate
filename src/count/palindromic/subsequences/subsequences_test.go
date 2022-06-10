package subsequences

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestSubsequences(t *testing.T) {
	suites := []suite{
		{
			"bccb",
			6,
		},
		{
			"abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba",
			104860361,
		},
		{
			"bddaabdbbccdcdcbbdbddccbaaccabbcacbadbdadbccddccdbdbdbdabdbddcccadddaaddbcbcbabdcaccaacabdbdaccbaacc",
			744991227,
		},
	}

	for _, s := range suites {
		ret := countPalindromicSubsequences(s.s)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
