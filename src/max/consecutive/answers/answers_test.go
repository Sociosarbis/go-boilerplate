package answers

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	answerKey string
	k         int
	ret       int
}

func TestAnswers(t *testing.T) {
	suites := []suite{
		{
			"TTFF",
			2,
			4,
		},
		{
			"TFFT",
			1,
			3,
		},
		{
			"TTFTTFTT",
			1,
			5,
		},
		{
			"TTTTTFTFFT",
			2,
			8,
		},
	}

	for _, s := range suites {
		ret := maxConsecutiveAnswers(s.answerKey, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
