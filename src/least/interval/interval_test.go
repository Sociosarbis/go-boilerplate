package interval

import (
	"testing"
)

type suite struct {
	tasks []byte
	n     int
	ret   int
}

func TestInterval(t *testing.T) {
	cases := []suite{
		{
			[]byte{65, 65, 65, 66, 66, 66},
			2,
			8,
		},
		{
			[]byte{65, 65, 65, 66, 66, 66},
			0,
			6,
		},
		{
			[]byte{65, 65, 65, 65, 65, 65, 66, 67, 68, 69, 70, 71},
			2,
			16,
		},
	}

	for _, c := range cases {
		ret := leastInterval(c.tasks, c.n)
		if ret != c.ret {
			t.Errorf("Failed, expected %v but get %v", c.ret, ret)
		}
	}
}
