package sort

func swap(arr *[]int, k int) {
	i := 0
	j := k - 1
	for i < j {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		i += 1
		j -= 1
	}
}

func pancakeSort(arr []int) []int {
	ret := []int{}
	for i := len(arr); i > 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if arr[j] == i {
				if j+1 != i {
					swap(&arr, j+1)
					ret = append(ret, j+1)
					swap(&arr, i)
					ret = append(ret, i)
				}
				break
			}
		}
	}
	return ret
}
