package operations

func countOperations(num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}
	var out int
	if num1 < num2 {
		num1, num2 = num2, num1
	}
	for num1 != 0 {
		diff := num1 - num2
		count := 1 + diff/num2
		out += count
		num1 -= num2 * count
		if num1 == 0 {
			return out
		}
		num1, num2 = num2, num1
	}
	return out
}
