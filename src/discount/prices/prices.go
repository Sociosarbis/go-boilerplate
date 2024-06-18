package prices

import (
	"fmt"
)

func isDigit(c byte) bool {
	return c >= 48 && c <= 57
}

func toInt(s string) int {
	base := 1
	var ret int
	for i := len(s) - 1; i >= 0; i-- {
		ret += int(s[i]-48) * base
		base *= 10
	}
	return ret
}

func discountPrices(sentence string, discount int) string {
	var start int
	n := len(sentence)
	ret := make([]byte, 0, n)
	fDiscount := 1 - float64(discount)/100
	var isPrice bool
	for i := 0; i <= n; i++ {
		if i == n || sentence[i] == ' ' {
			if i > start {
				if isPrice && start+1 < i {
					ret = append(ret, '$')
					price := float64(toInt(sentence[start+1:i])) * fDiscount
					ret = append(ret, []byte(fmt.Sprintf("%.2f", price))...)
				} else {
					ret = append(ret, []byte(sentence[start:i])...)
				}
				if i != n {
					ret = append(ret, ' ')
				}
			}
			isPrice = false
			start = i + 1
		} else if isPrice {
			if !isDigit(sentence[i]) {
				isPrice = false
			}
		} else {
			if i == start && sentence[i] == '$' {
				isPrice = true
			}
		}
	}
	return string(ret)
}
