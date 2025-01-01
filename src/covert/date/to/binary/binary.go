package binary

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	parts := strings.Split(date, "-")
	for i, part := range parts {
		num, _ := strconv.ParseInt(part, 10, 32)
		parts[i] = strconv.FormatInt(num, 2)
	}
	return strings.Join(parts, "-")
}
