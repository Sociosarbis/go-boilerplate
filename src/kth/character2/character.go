package character2

func getPath(k int64, path *[]int64) {
	*path = append(*path, k)
	var i int
	for 1<<i < k {
		i++
	}
	if i > 0 {
		getPath(k-(1<<(i-1)), path)
	}
}

func kthCharacter(k int64, operations []int) byte {
	path := make([]int64, 0, len(operations))
	getPath(k, &path)
	n := len(path)
	var out byte = 'a'
	var j int
	for i := n - 2; i >= 0; i-- {
		index := path[i]
		for 1<<j < index {
			j++
		}
		op := operations[j-1]
		if op == 1 {
			out = (out-'a'+1)%26 + 'a'
		}
	}
	return out
}
