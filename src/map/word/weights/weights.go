package weights

const mod int = 26

func mapWordWeights(words []string, weights []int) string {
	out := make([]byte, 0, len(words))
	for _, word := range words {
		var temp int
		for _, c := range word {
			temp = (temp + weights[c-'a']) % mod
		}
		out = append(out, byte('z'-temp))
	}
	return string(out)
}
