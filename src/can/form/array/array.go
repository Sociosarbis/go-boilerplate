package array

func canFormArray(arr []int, pieces [][]int) bool {
	indices := make(map[int]int, len(arr))
	for i, num := range arr {
		indices[num] = i
	}
	for _, piece := range pieces {
		if index, ok := indices[piece[0]]; ok {
			for i := 1; i < len(piece); i++ {
				if index+i >= len(arr) || arr[index+i] != piece[i] {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}
