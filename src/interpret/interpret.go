package interpret

func interpret(command string) string {
	ret := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			ret += "G"
		} else {
			if command[i+1] == ')' {
				ret += "o"
				i++
			} else {
				ret += "al"
				i += 3
			}
		}
	}
	return ret
}
