package operations

func finalValueAfterOperations(operations []string) int {
	ret := 0
	for _, operation := range operations {
		if operation == "++X" || operation == "X++" {
			ret++
		} else {
			ret--
		}
	}
	return ret
}
