package evaluate

import "strconv"

func evaluate(expression string) int {
	i := 0
	return dfs(expression, map[string]int{}, &i)
}

func readToken(expression string, i *int) string {
	l := *i
	for *i < len(expression) && (expression[*i] != ' ' && expression[*i] != ')') {
		*i++
	}
	return expression[l:*i]
}

func readValue(t string, ctx map[string]int) int {
	if num, ok := ctx[t]; ok {
		return num
	} else {
		num, _ := strconv.ParseInt(t, 10, 32)
		return int(num)
	}
}

func dfs(expression string, ctx map[string]int, i *int) int {
	if *i < len(expression) && expression[*i] == '(' {
		*i++
	}
	op := readToken(expression, i)
	var name string
	if op == "let" {
		for *i < len(expression) && expression[*i] != ')' {
			if expression[*i] != ' ' {
				if expression[*i] != '(' {
					t := readToken(expression, i)
					if len(name) != 0 {
						ctx[name] = readValue(t, ctx)
						name = name[:0]
					} else {
						name = t
					}
				} else {
					nextCtx := make(map[string]int, len(ctx))
					for k, v := range ctx {
						nextCtx[k] = v
					}
					ret := dfs(expression, nextCtx, i)
					*i++
					if len(name) != 0 {
						ctx[name] = ret
						name = name[:0]
					} else {
						return ret
					}
				}
			} else {
				*i++
			}
		}
		if len(name) != 0 {
			return readValue(name, ctx)
		}
	} else if op == "mult" || op == "add" {
		var left int
		j := 0
		for *i < len(expression) && expression[*i] != ')' {
			if expression[*i] != ' ' {
				if expression[*i] != '(' {
					t := readToken(expression, i)
					if j == 0 {
						left = readValue(t, ctx)
						j++
					} else {
						if op == "add" {
							return left + readValue(t, ctx)
						} else {
							return left * readValue(t, ctx)
						}
					}
				} else {
					nextCtx := make(map[string]int, len(ctx))
					for k, v := range ctx {
						nextCtx[k] = v
					}
					ret := dfs(expression, nextCtx, i)
					*i++
					if j == 0 {
						left = ret
						j++
					} else {
						if op == "add" {
							return left + ret
						} else {
							return left * ret
						}
					}
				}
			} else {
				*i++
			}
		}
	}
	num, _ := strconv.ParseInt(op, 10, 32)
	return int(num)
}
