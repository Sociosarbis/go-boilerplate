package valid

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	code string
	ret  bool
}

func TestValid(t *testing.T) {
	suites := []suite{
		{
			"<DIV>This is the first line <![CDATA[<div>]]></DIV>",
			true,
		},
		{
			"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>",
			true,
		},
		{
			"<A>  <B> </A>   </B>",
			false,
		},
		{
			"<A><></A>",
			false,
		},
	}

	for _, s := range suites {
		ret := isValid(s.code)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
