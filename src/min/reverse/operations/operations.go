package operations

import "github.com/boilerplate/src/tree/avl"

func minReverseOperations(n int, p int, banned []int, k int) []int {
	ret := make([]int, n)
	for _, i := range banned {
		ret[i] = -1
	}
	bfs := []int{p}
	qs := make([]*avl.AvlNode[avl.Integer], 2)
	for i := 0; i < n; i++ {
		if ret[i] != -1 {
			qs[i&1] = qs[i&1].Insert(avl.Integer(i))
		}
	}
	l := len(bfs)
	t := 1
	for l != 0 {
		for i := 0; i < l; i++ {
			end := k - 1
			if n-1-bfs[i] < end {
				end = (n-1-bfs[i])*2 - end
			}
			start := -k + 1
			if start+bfs[i] < 0 {
				start = -(bfs[i]*2 + start)
			}
			start, end = bfs[i]+start, bfs[i]+end
			toRemove := []int{}
			q := qs[start&1]
			iter := q.LowerBound(avl.Integer(start))
			cur := iter.GetCur()
			for ; cur != nil && int(cur.GetValue()) <= end; cur = iter.Next() {
				index := int(cur.GetValue())
				toRemove = append(toRemove, index)
				if index < 0 || index >= n || index == p || ret[index] != 0 {
					continue
				}
				ret[index] = t
				bfs = append(bfs, index)
			}
			for _, v := range toRemove {
				q = q.Delete(avl.Integer(v))
			}
			qs[start&1] = q
		}
		t++
		bfs = bfs[l:]
		l = len(bfs)
	}
	for i := 0; i < n; i++ {
		if i != p && ret[i] == 0 {
			ret[i] = -1
		}
	}
	return ret
}
