package sorted

func isAlienSorted(words []string, order string) bool {
	priorityList := make([]int, 26)
	for i, c := range order {
		priorityList[c-97] = i
	}

	for i := 0; i < len(words)-1; i++ {
		left := words[i]
		right := words[i+1]

		j := 0

		valid := false
		for j < len(left) && j < len(right) {
			if priorityList[left[j]-97] > priorityList[right[j]-97] {
				return false
			} else if priorityList[left[j]-97] < priorityList[right[j]-97] {
				valid = true
				break
			}
			j++
		}

		if !valid && len(left) > len(right) {
			return false
		}
	}

	return true
}
