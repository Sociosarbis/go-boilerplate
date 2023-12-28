package list

func Insert[T any](arr []T, index int, value T) []T {
	var item T
	arr = append(arr, item)
	copy(arr[index+1:], arr[index:])
	arr[index] = value
	return arr
}

func Remove[T any](arr []T, index int) ([]T, T) {
	var removed T
	if index < len(arr) {
		removed = arr[index]
		copy(arr[index:], arr[index+1:])
		arr = arr[:len(arr)-1]
	}
	return arr, removed
}
