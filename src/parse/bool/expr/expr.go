package expr

func dfs(expression string, i int) (bool, int) {
	if expression[i] == 't' {
		return true, i + 1
	} else if expression[i] == 'f' {
		return false, i + 1
	} else {
		isAnd := true
		isNot := false
		ret := true
		if expression[i] == '|' {
			isAnd = false
			ret = false
		} else if expression[i] == '!' {
			isNot = true
		}
		i += 2
		for {
			value, index := dfs(expression, i)
			if isAnd {
				if ret && !value {
					ret = false
				}
			} else {
				if !ret && value {
					ret = true
				}
			}
			if expression[index] == ')' {
				if isNot {
					return !ret, index + 1
				} else {
					return ret, index + 1
				}
			}
			i = index + 1
		}
	}
}

func parseBoolExpr(expression string) bool {
	ret, _ := dfs(expression, 0)
	return ret
}
