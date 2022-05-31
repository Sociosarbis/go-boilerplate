package order

func alienOrder(words []string) string {
	degreeZeroes := []int{}

	degrees := make([]map[byte]bool, 26)

	for i := 1; i < len(words); i++ {
		j := 0
		for j < len(words[i-1]) && j < len(words[i]) {
			if words[i-1][j] != words[i][j] {
				k1 := words[i-1][j] - 97
				k2 := words[i][j] - 97
				if degrees[k2] == nil {
					degrees[k2] = map[byte]bool{}
				}
				degrees[k2][k1] = true
				break
			}
			j++
		}
		if (j == len(words[i-1]) || j == len(words[i])) && len(words[i-1]) > len(words[i]) {
			return ""
		}
	}

	ret := []byte{}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if degrees[words[i][j]-97] == nil {
				degrees[words[i][j]-97] = make(map[byte]bool, 0)
			}
		}
	}

	for i := 0; i < 26; i++ {
		if degrees[i] != nil && len(degrees[i]) == 0 {
			degreeZeroes = append(degreeZeroes, i)
		}
	}

	for len(degreeZeroes) != 0 {
		n := len(degreeZeroes)
		for i := 0; i < n; i++ {
			c := byte(degreeZeroes[i])
			ret = append(ret, c+97)
			for j := 0; j < 26; j++ {
				if degrees[j] != nil {
					if _, ok := degrees[j][c]; ok {
						delete(degrees[j], c)
						if len(degrees[j]) == 0 {
							degreeZeroes = append(degreeZeroes, j)
						}
					}
				}
			}
		}
		degreeZeroes = degreeZeroes[n:]
	}

	for j := 0; j < 26; j++ {
		if degrees[j] != nil && len(degrees[j]) != 0 {
			return ""
		}
	}

	return string(ret)
}
