package path

func lengthLongestPath(input string) int {
	stack := []int{}
	temp := 0
	ret := 0
	i := 0
	for i < len(input) {
		j := i
		hasDot := false
		for j < len(input) && input[j] != '\n' {
			if !hasDot && input[j] == '.' {
				hasDot = true
			}
			j++
		}
		length := j - i
		if hasDot {
			if temp+length+len(stack) > ret {
				ret = temp + length + len(stack)
			}
		} else {
			temp += length
			stack = append(stack, length)
		}

		j++
		i = j
		for j < len(input) && input[j] == '\t' {
			j++
		}
		for k := j - i; k < len(stack); k++ {
			temp -= stack[k]
		}
		stack = stack[0 : j-i]
		i = j
	}
	return ret
}
