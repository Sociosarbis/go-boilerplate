package message

func decodeMessage(key string, message string) string {
	met := make([]bool, 26)
	dict := make([]byte, 26)
	var i byte = 97
	for _, k := range key {
		if k != ' ' && !met[k-97] {
			dict[k-97] = byte(i)
			met[k-97] = true
			i++
		}
	}
	ret := []byte(message)
	for i, k := range ret {
		if k != ' ' {
			ret[i] = dict[k-97]
		}
	}
	return string(ret)
}
