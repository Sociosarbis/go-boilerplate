package path

import (
	"strconv"
	"strings"
)

type q struct {
	isUpdate bool
	u        int
	v        int
	b        byte
}

type bit struct {
	tree []int
	n    int
}

func lowBit(index int) int {
	return index & -index
}

func newBit(n int) bit {
	return bit{
		tree: make([]int, n+1),
		n:    n,
	}
}

func (self *bit) update(x, v int) {
	for x <= self.n {
		self.tree[x] ^= v
		x += lowBit(x)
	}
}

func (self *bit) query(x int) int {
	var out int
	for x > 0 {
		out ^= self.tree[x]
		x -= lowBit(x)
	}
	return out
}

func (self *bit) updateRange(l, r, v int) {
	self.update(l, v)
	self.update(r+1, v)
}

func newQ(query string) q {
	i := strings.IndexByte(query, ' ')
	ni := i + 1 + strings.IndexByte(query[i+1:], ' ')
	u, _ := strconv.ParseInt(query[i+1:ni], 10, 32)
	var out q
	out.u = int(u)
	if query[:i] == "update" {
		out.isUpdate = true
		i = ni
		out.b = query[i+1]
	} else {
		i = ni
		v, _ := strconv.ParseInt(query[i+1:], 10, 32)
		out.v = int(v)
	}
	return out
}

func getLog(n int) int {
	var out int
	for n != 0 {
		n >>= 1
		out++
	}
	return out
}

func dfs(g, fa [][]int, l, r, dep []int, timer *int, u, p int) {
	*timer++
	l[u] = *timer
	fa[u][0] = p
	for _, v := range g[u] {
		if v == p {
			continue
		}
		dep[v] = dep[u] + 1
		dfs(g, fa, l, r, dep, timer, v, u)
	}
	r[u] = *timer
}

func lca(fa [][]int, dep []int, log, u, v int) int {
	if dep[u] < dep[v] {
		u, v = v, u
	}
	diff := dep[u] - dep[v]
	for i := 0; i < log; i++ {
		if (diff>>i)&1 == 1 {
			u = fa[u][i]
		}
	}
	// 相同深度，但是不相等
	if u == v {
		return u
	}
	for i := log - 1; i >= 0; i-- {
		if fa[u][i] != fa[v][i] {
			u, v = fa[u][i], fa[v][i]
		}
	}
	return fa[u][0]
}

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
	out := make([]bool, 0, len(queries))
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	log := getLog(n)
	// 距离2^i的祖先节点
	fa := make([][]int, n)
	for i := 0; i < n; i++ {
		fa[i] = make([]int, log)
	}
	dep, l, r := make([]int, n), make([]int, n), make([]int, n)
	var timer int
	dfs(g, fa, l, r, dep, &timer, 0, 0)

	for j := 1; j < log; j++ {
		for i := 0; i < n; i++ {
			fa[i][j] = fa[fa[i][j-1]][j-1]
		}
	}
	b := newBit(n + 1)
	for i := 0; i < n; i++ {
		// 获得根节点到子节点路径的异或
		b.updateRange(l[i], r[i], 1<<(s[i]-'a'))
	}
	bs := []byte(s)
	for _, s := range queries {
		query := newQ(s)
		if query.isUpdate {
			b.updateRange(l[query.u], r[query.u], 1<<(bs[query.u]-'a')^1<<(query.b-'a'))
			bs[query.u] = query.b
		} else {
			w := lca(fa, dep, log, query.u, query.v)
			mask := b.query(l[query.u]) ^ b.query(l[query.v]) ^ (1 << (bs[w] - 'a'))
			out = append(out, mask == 0 || mask&(mask-1) == 0)
		}
	}
	return out
}
