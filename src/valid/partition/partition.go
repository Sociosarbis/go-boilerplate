package partition

var bitOffset = [3]int{0, 2, 5}

var initBits byte = set(set(set(0, 1, 0), 1, 1), 1, 2)
var fullBits byte = set(set(set(0, 2, 0), 3, 1), 3, 2)

func has(bits byte, i byte, level byte) bool {
	var mask byte = (1 << (i - 1)) << bitOffset[level]
	return (bits & mask) == mask
}

func set(bits byte, i byte, level byte) byte {
	return bits | ((1 << (i - 1)) << bitOffset[level])
}

func validPartition(nums []int) bool {
	n := len(nums)

	prevStatus := initBits

	for i := 1; i < n; i++ {
		num := nums[i]
		prevNum := nums[i-1]
		var status byte
		if (prevStatus & fullBits) != 0 {
			status = initBits
		}
		if num == prevNum {
			if has(prevStatus, 1, 0) {
				status = set(status, 2, 0)
			}
			var j byte
			for j = 1; j < 3; j++ {
				if has(prevStatus, j, 1) {
					status = set(status, j+1, 1)
				}
			}
		} else if num == prevNum+1 {
			var j byte = 1
			for ; j < 3; j++ {
				if has(prevStatus, j, 2) {
					status = set(status, j+1, 2)
				}
			}
		}
		if status == 0 {
			return false
		}
		prevStatus = status
	}

	return (prevStatus & fullBits) != 0
}
