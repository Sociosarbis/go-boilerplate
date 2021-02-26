package words

func findNumOfValidWords(words []string, puzzles []string) []int {
	wordsBits := map[int]int{}

	ret := make([]int, len(puzzles))

	for _, word := range words {
		num := 0
		for _, c := range word {
			num |= 1 << (c - 97)
		}
		count, ok := wordsBits[num]
		if !ok {
			wordsBits[num] = 1
		} else {
			wordsBits[num] = count + 1
		}
	}

	for i, puzzle := range puzzles {
		num := 0
		firstBit := 1 << (puzzle[0] - 97)
		for _, c := range puzzle {
			num |= 1 << (c - 97)
		}
		/** 从大到小遍历子mask的方法，原因是j - 1是下个最大整数，然后and操作以后得到下个最大的子mask */
		for j := num; j != 0; j = (j - 1) & num {
			if j&firstBit != 0 {
				count, ok := wordsBits[j]
				if ok {
					ret[i] += count
				}
			}
		}
	}
	return ret
}
