package similarity

type Item struct {
	s string
	i int
}

func kSimilarity(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	}
	item := Item{
		s: s1,
		i: len(s2),
	}
	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			item.i = i
			break
		}
	}
	ret := 0
	bfs := []Item{item}
	for len(bfs) != 0 {
		ret++
		n := len(bfs)
		for i := 0; i < n; i++ {
			bs := []byte(bfs[i].s)
			index := bfs[i].i
			for j := bfs[i].i; j < len(bs); j++ {
				if bs[j] == s2[index] {
					bs[j], bs[index] = bs[index], bs[j]
					item := Item{
						s: string(bs),
						i: len(s2),
					}
					for k := bfs[i].i + 1; k < len(s2); k++ {
						if item.s[k] != s2[k] {
							item.i = k
							break
						}
					}
					bs[j], bs[index] = bs[index], bs[j]
					if item.i < len(s2) {
						bfs = append(bfs, item)
					} else {
						return ret
					}
				}
			}
		}
		bfs = bfs[n:]
	}
	return ret
}
