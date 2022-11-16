package permutation

func isIdealPermutation(nums []int) bool {
	arr := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if len(arr) > 1 {
			if arr[0] > nums[i] {
				return false
			}
		}
		if arr[len(arr)-1] > nums[i] && arr[len(arr)-1] != nums[i-1] {
			return false
		}
		j := 0
		for ; j < len(arr); j++ {
			if arr[j] >= nums[i] {
				break
			}
		}
		if j > 0 || len(arr) < 2 {
			arr = append(arr, 0)
			copy(arr[j+1:], arr[j:])
			arr[j] = nums[i]
			if len(arr) > 2 {
				arr = arr[1:]
			}
		}
	}
	return true
}
