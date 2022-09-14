package mean

import (
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)

	count := len(arr) / 20
	temp := 0
	for i := count; i+count < len(arr); i++ {
		temp += arr[i]
	}
	return float64(temp) / float64(len(arr)-count*2)
}
