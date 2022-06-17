package zeros

func duplicateZeros(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count++
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if i+count < len(arr) {
				arr[i+count] = 0
			}
			count--
		}
		if count > 0 && i+count < len(arr) {
			arr[i+count] = arr[i]
		}
	}
}
