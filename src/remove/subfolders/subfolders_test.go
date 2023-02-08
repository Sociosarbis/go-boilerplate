package subfolders

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	folder []string
	ret    []string
}

func TestSubfolders(t *testing.T) {
	suites := []suite{
		{
			[]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			[]string{"/a", "/c/d", "/c/f"},
		},
		{
			[]string{"/a", "/a/b/c", "/a/b/d"},
			[]string{"/a"},
		},
		{
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			[]string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
	}

	for _, s := range suites {
		ret := removeSubfolders(s.folder)
		m1 := map[string]bool{}
		for _, folder := range s.ret {
			m1[folder] = true
		}
		m2 := map[string]bool{}
		for _, folder := range ret {
			m2[folder] = true
		}
		if !reflect.DeepEqual(m1, m2) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
