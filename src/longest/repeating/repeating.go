package repeating

type node struct {
	leftCount  int
	rightCount int
	max        int
	leftChar   byte
	rightChar  byte
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func split(i, l, r int) (int, int, int, int, int, int) {
	li, ri := i*2+1, i*2+2
	mid := (l + r) / 2
	ll, lr, rl, rr := l, mid, mid+1, r
	return li, ll, lr, ri, rl, rr
}

func buildTree(nodes []node, s string, i, l, r int) {
	if l == r {
		nodes[i].leftCount, nodes[i].rightCount, nodes[i].max, nodes[i].leftChar, nodes[i].rightChar = 1, 1, 1, s[l], s[l]
	} else {
		li, ll, lr, ri, rl, rr := split(i, l, r)
		buildTree(nodes, s, li, ll, lr)
		buildTree(nodes, s, ri, rl, rr)
		bottomUp(nodes, i, l, r)
	}
}

func bottomUp(nodes []node, i, l, r int) {
	li, ll, lr, ri, rl, rr := split(i, l, r)
	mx := max(nodes[li].max, nodes[ri].max)
	nodes[i].leftCount, nodes[i].rightCount, nodes[i].max, nodes[i].leftChar, nodes[i].rightChar = nodes[li].leftCount, nodes[ri].rightCount, mx, nodes[li].leftChar, nodes[ri].rightChar
	if nodes[li].rightChar == nodes[ri].leftChar {
		count := nodes[li].rightCount + nodes[ri].leftCount
		if count > mx {
			nodes[i].max = count
		}
		if lr-ll+1 == nodes[li].leftCount {
			nodes[i].leftCount = count
		}
		if rr-rl+1 == nodes[ri].rightCount {
			nodes[i].rightCount = count
		}
	}
}

func update(nodes []node, i, l, r, index int, c byte) {
	if index < l || index > r {
		return
	}
	if l == r {
		nodes[i].leftCount, nodes[i].rightCount, nodes[i].max, nodes[i].leftChar, nodes[i].rightChar = 1, 1, 1, c, c
	} else {
		li, ll, lr, ri, rl, rr := split(i, l, r)
		update(nodes, li, ll, lr, index, c)
		update(nodes, ri, rl, rr, index, c)
		bottomUp(nodes, i, l, r)
	}
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	m := len(s)
	count := 1
	for m != 0 {
		m >>= 1
		count++
	}
	nodes := make([]node, 1<<count)
	m = len(s)
	buildTree(nodes, s, 0, 0, m-1)
	out := make([]int, len(queryCharacters))
	for i, index := range queryIndices {
		update(nodes, 0, 0, m-1, index, queryCharacters[i])
		out[i] = nodes[0].max
	}
	return out
}
