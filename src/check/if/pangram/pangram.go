package pangram

func checkIfPangram(sentence string) bool {
	m := make([]bool, 26)
	temp := 0
	for _, c := range sentence {
		if !m[c-97] {
			m[c-97] = true
			temp++
		}
	}
	return temp == 26
}
