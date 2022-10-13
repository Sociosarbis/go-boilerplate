package sorted

type chunk struct {
	min int
	max int
}

func maxChunksToSorted(arr []int) int {
	chunks := []chunk{
		{
			arr[0],
			arr[0],
		},
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] > chunks[len(chunks)-1].max {
			chunks = append(chunks, chunk{
				arr[i],
				arr[i],
			})
		} else {
			j := len(chunks) - 1
			if arr[i] < chunks[j].min {
				chunks[j].min = arr[i]
			}
			for j > 0 {
				if chunks[j].min < chunks[j-1].max {
					chunks[j-1].max = chunks[j].max
					if chunks[j].min < chunks[j-1].min {
						chunks[j-1].min = chunks[j].min
					}
					chunks = chunks[:j]
					j--
				} else {
					break
				}
			}
		}
	}
	return len(chunks)
}
