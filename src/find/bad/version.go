package findBadVersion

var badVersion = 0

func isBadVersion(version int) bool {
	return version >= badVersion
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	ret := n
	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			ret = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ret
}
