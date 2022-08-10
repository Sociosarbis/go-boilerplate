package equation

import (
	"fmt"
	"strconv"
)

func readItem(equation string, i int) (int, uint8, int) {
	if equation[i] == '=' {
		return 0, 3, i + 1
	} else {
		l := i
		i++
		for i < len(equation) && ((equation[i] >= 48 && equation[i] <= 57) || equation[i] == 'x') {
			i++
		}
		text := equation[l:i]
		if text[len(text)-1] == 'x' {
			var v int
			num := text[:len(text)-1]
			if len(num) > 1 || (len(num) > 0 && num[0] >= 48 && num[0] <= 57) {
				res, _ := strconv.ParseInt(num, 10, 32)
				v = int(res)
			} else {
				if len(num) == 0 || num[0] == '+' {
					v = 1
				} else {
					v = -1
				}
			}
			return v,
				2, i
		} else {
			v, _ := strconv.ParseInt(text, 10, 32)
			return int(v), 1, i
		}
	}
}

func solveEquation(equation string) string {
	i := 0
	l := 0
	r := 0
	reversed := false
	for i < len(equation) {
		value, type_, nextI := readItem(equation, i)
		if type_ == 3 {
			reversed = true
		} else if type_ == 1 {
			if reversed {
				r += value
			} else {
				r -= value
			}
		} else {
			if reversed {
				l -= value
			} else {
				l += value
			}
		}
		i = nextI
	}
	if l == 0 {
		if r == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		return fmt.Sprintf("x=%v", r/l)
	}
}
