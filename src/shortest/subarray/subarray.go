package subarray

type item struct {
	i int
	v int64
}

func lessThen(items []item, v int64) int {
	l := 0
	r := len(items) - 1
	for l <= r {
		mid := (l + r) / 2
		if items[mid].v <= v {
			l = mid + 1
		} else {
			if mid > 0 && items[mid-1].v > v {
				r = mid - 1
			} else {
				return mid
			}
		}
	}
	return l
}

func shortestSubarray(nums []int, k int) int {
	arr := []item{}
	target := int64(k)
	var temp int64 = 0
	v := -1
	for i, num := range nums {
		temp += int64(num)
		n := lessThen(arr, temp-target)
		if temp >= target {
			if v == -1 || i+1 < v {
				v = i + 1
			}
		}
		for j := 0; j < n; j++ {
			// 只跟差值大于k对比
			if v == -1 || i-arr[j].i < v {
				v = i - arr[j].i
			}
		}
		// 对比过的，不用再对比，因为后面即便符合条件，也不可能更小
		arr = arr[n:]
		index := lessThen(arr, temp)
		if index >= len(arr) {
			arr = append(arr, item{i, temp})
		} else {
			// 大于它的可以截断
			arr = arr[:index]
			arr = append(arr, item{})
			copy(arr[index+1:], arr[index:])
			arr[index] = item{i, temp}
		}
	}
	return v
}
