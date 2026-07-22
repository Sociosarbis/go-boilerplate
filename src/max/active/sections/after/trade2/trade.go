package trade2

type segment struct {
	i      int
	l      int
	r      int
	isZero bool
}

type node struct {
	l int
	r int
	s int
	e int
	v int
}

func buildTree(nodes []node, segments []segment, i, l, r int) {
	if l == r {
		var v int
		s, e := segments[l].l, segments[l].r
		if !segments[l].isZero {
			if l > 0 && l+1 < len(segments) {
				v = segments[l+1].r + segments[l-1].r - segments[l+1].l - segments[l-1].l + 2
				s, e = segments[l-1].l, segments[l+1].r
			}
		}
		nodes[i] = node{
			l,
			r,
			s,
			e,
			v,
		}
		return
	}
	left := i*2 + 1
	right := left + 1
	mid := (l + r) / 2
	buildTree(nodes, segments, left, l, mid)
	buildTree(nodes, segments, right, mid+1, r)
	v := nodes[left].v
	if nodes[right].v > v {
		v = nodes[right].v
	}
	nodes[i] = node{
		l: l,
		r: r,
		s: nodes[left].s,
		e: nodes[right].e,
		v: v,
	}
}

func count(s segment, l, r int) int {
	if s.l > l {
		l = s.l
	}
	if s.r < r {
		r = s.r
	}
	if l > r {
		return 0
	}
	return r - l + 1
}

func search(nodes []node, segments []segment, i, l, r int) int {
	sl, sr := nodes[i].l, nodes[i].r
	rl, rr := nodes[i].s, nodes[i].e
	if rl >= l && rr <= r {
		return nodes[i].v
	}
	if rr < l || rl > r {
		return 0
	}
	if sl == sr {
		if nodes[i].v == 0 {
			return 0
		}
		var c1, c3 int
		if sl > 0 {
			c1 = count(segments[sl-1], l, r)
		}
		if sl+1 < len(segments) {
			c3 = count(segments[sl+1], l, r)
		}
		if c1 == 0 || c3 == 0 {
			return 0
		}
		return c1 + c3
	}
	left := i*2 + 1
	right := left + 1
	v1 := search(nodes, segments, left, l, r)
	v2 := search(nodes, segments, right, l, r)
	if v1 > v2 {
		return v1
	}
	return v2
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	segments := []segment{}
	var start int
	n := len(s)
	var base int
	if s[0] == '1' {
		base = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			base++
		}
		if s[i] != s[i-1] {
			segments = append(segments, segment{
				i:      len(segments),
				l:      start,
				r:      i - 1,
				isZero: s[start] == '0',
			})
			start = i
		}
	}
	segments = append(segments, segment{
		i:      len(segments),
		l:      start,
		r:      n - 1,
		isZero: s[start] == '0',
	})
	count := 1
	m := len(segments)
	for m != 0 {
		m >>= 1
		count++
	}
	nodes := make([]node, 1<<count)
	buildTree(nodes, segments, 0, 0, len(segments)-1)
	out := make([]int, len(queries))
	for i, query := range queries {
		out[i] = base + search(nodes, segments, 0, query[0], query[1])
	}
	return out
}
