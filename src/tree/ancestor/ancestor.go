package ancestor

type TreeAncestor struct {
	parent []int
	levels []int
}

func Constructor(n int, parent []int) TreeAncestor {
	ancestorPaths := make([][]int, n)
	ancestorPaths[0] = []int{0}
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		if children[parent[i]] == nil {
			children[parent[i]] = []int{i}
		} else {
			children[parent[i]] = append(children[parent[i]], i)
		}
	}
	levels := make([]int, n)
	for i := range levels {
		levels[i] = n
	}
	treeAncestor := TreeAncestor{
		parent,
		levels,
	}
	treeAncestor.buildLevels(0, -1, 0, children)
	return treeAncestor
}

func (this *TreeAncestor) buildLevels(node int, prevLevelNode int, layer int, children [][]int) {
	if layer%100 == 0 {
		this.levels[node] = prevLevelNode
		prevLevelNode = node
	}
	for _, child := range children[node] {
		this.buildLevels(child, prevLevelNode, layer+1, children)
	}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for node != -1 && this.levels[node] == len(this.levels) && k != 0 {
		node = this.parent[node]
		k -= 1
	}
	for node != -1 && k >= 100 && this.levels[node] != len(this.levels) {
		node = this.levels[node]
		k -= 100
	}
	for node != -1 && k != 0 {
		node = this.parent[node]
		k -= 1
	}
	return node
}
