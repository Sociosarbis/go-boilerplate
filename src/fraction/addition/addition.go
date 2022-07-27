package addition

import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string {
	i := 0
	x := readNum(expression, &i)
	i++
	y := readNum(expression, &i)
	for i < len(expression) {
		op := expression[i]
		i++
		tempX := readNum(expression, &i)
		i++
		tempY := readNum(expression, &i)
		nextY := tempY * y
		if op == '+' {
			x = x*tempY + tempX*y
		} else {
			x = x*tempY - tempX*y
		}
		x, y = gcd(x, nextY)
	}
	x, y = gcd(x, y)
	return fmt.Sprintf("%v/%v", x, y)
}

func gcd(x, y int) (int, int) {
	if x == 0 {
		return x, 1
	}
	tempX := x
	if tempX < 0 {
		tempX = -tempX
	}
	tempY := y
	if tempX > tempY {
		tempX, tempY = tempY, tempX
	}
	for tempY%tempX != 0 {
		tempX, tempY = tempY%tempX, tempX
	}
	return x / tempX, y / tempX
}

func readNum(expression string, i *int) int {
	start := *i
	if expression[*i] == '-' {
		*i++
	}
	for *i < len(expression) && expression[*i] >= 48 && expression[*i] <= 57 {
		*i++
	}
	ret, _ := strconv.ParseInt(expression[start:*i], 10, 32)
	return int(ret)
}
