package avl

import (
	"testing"
)

func TestAvlTree(t *testing.T) {
	var root *AvlNode[Integer]
	for i := 0; i < 10; i++ {
		root = root.Insert(Integer(i))
	}
	root.Delete(Integer(5))
	iter := root.LowerBound(5)
	cur := iter.GetCur()
	if cur != nil {
		for ; cur != nil; cur = iter.Next() {
			t.Logf("%d", cur.GetValue())
		}
	}
}
