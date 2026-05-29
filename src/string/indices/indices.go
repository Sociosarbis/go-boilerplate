package indices

type node struct {
	next [26]*node
	i    int
}

func add(wordsContainer []string, index int, nd *node) {
	word := wordsContainer[index]
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		if nd.next[word[i]-'a'] == nil {
			nd.next[word[i]-'a'] = &node{
				i: index,
			}
		} else if len(wordsContainer[nd.next[word[i]-'a'].i]) > n {
			nd.next[word[i]-'a'].i = index
		}
		nd = nd.next[word[i]-'a']
	}
}

func ans(nd *node, word string) int {
	n := len(word)
	out := -1
	for i := n - 1; i >= 0; i-- {
		if nd.next[word[i]-'a'] != nil {
			out = nd.next[word[i]-'a'].i
			nd = nd.next[word[i]-'a']
		} else {
			break
		}
	}
	return out
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := &node{
		i: -1,
	}
	n := len(wordsContainer)
	var minIndex int
	for i := 0; i < n; i++ {
		add(wordsContainer, i, root)
		if len(wordsContainer[i]) < len(wordsContainer[minIndex]) {
			minIndex = i
		}
	}
	n = len(wordsQuery)
	out := make([]int, n)
	for i := 0; i < n; i++ {
		temp := ans(root, wordsQuery[i])
		if temp != -1 {
			out[i] = temp
		} else {
			out[i] = minIndex
		}
	}
	return out
}
