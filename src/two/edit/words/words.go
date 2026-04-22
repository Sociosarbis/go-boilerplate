package words

func check(a, b string) bool {
	var count int
	for i, c := range a {
		if byte(c) == b[i] {
			continue
		}
		count++
		if count > 2 {
			return false
		}
	}
	return true
}

func twoEditWords(queries []string, dictionary []string) []string {
	out := make([]string, 0, len(queries))
	for _, query := range queries {
		for _, word := range dictionary {
			if check(query, word) {
				out = append(out, query)
				break
			}
		}
	}
	return out
}
