package sizes

type node struct {
	children []*node
	i        int
}

func dfs(node *node, parent, bToIndex []int, s string) {
	b := s[node.i] - 'a'
	oldIndex := bToIndex[b]
	if oldIndex != -1 {
		parent[node.i] = oldIndex
	}
	bToIndex[b] = node.i
	for _, child := range node.children {
		dfs(child, parent, bToIndex, s)
	}
	bToIndex[b] = oldIndex
}

func dfsSize(node *node, out []int) int {
	out[node.i] = 1
	for _, child := range node.children {
		out[node.i] += dfsSize(child, out)
	}
	return out[node.i]
}

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	nodes := make([]node, n)
	for i := 0; i < n; i++ {
		nodes[i] = node{
			i: i,
		}
	}
	for i := 1; i < n; i++ {
		nodes[parent[i]].children = append(nodes[parent[i]].children, &nodes[i])
	}
	bToIndex := make([]int, 26)
	for i := 0; i < 26; i++ {
		bToIndex[i] = -1
	}
	dfs(&nodes[0], parent, bToIndex, s)
	for i := 0; i < n; i++ {
		nodes[i].children = nil
	}
	for i := 1; i < n; i++ {
		nodes[parent[i]].children = append(nodes[parent[i]].children, &nodes[i])
	}
	out := make([]int, n)
	dfsSize(&nodes[0], out)
	return out
}
