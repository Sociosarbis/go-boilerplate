package latin

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	sentence string
	ret      string
}

func TestLatin(t *testing.T) {
	suites := []suite{
		{
			"I speak Goat Latin",
			"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			"The quick brown fox jumped over the lazy dog",
			"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}

	for _, s := range suites {
		ret := toGoatLatin(s.sentence)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
