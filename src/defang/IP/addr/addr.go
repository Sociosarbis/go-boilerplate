package addr

func defangIPaddr(address string) string {
	ret := []byte{}
	for i := 0; i < len(address); i++ {
		if address[i] != '.' {
			ret = append(ret, address[i])
		} else {
			ret = append(ret, '[', '.', ']')
		}
	}
	return string(ret)
}
