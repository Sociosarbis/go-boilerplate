package elements

func abs(x, y int) int {
	if y >= x {
		return y - x
	}
	return x - y
}

func findClosestElements(arr []int, k int, x int) []int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) >> 1
		if arr[mid] < x {
			if mid+1 < len(arr) && (abs(arr[mid+1], x) < x-arr[mid] || arr[mid+1] == arr[mid]) {
				l = mid + 1
				continue
			}
		} else if arr[mid] > x {
			if mid > 0 && abs(arr[mid-1], x) <= arr[mid]-x {
				r = mid - 1
				continue
			}
		}
		l = mid
		break
	}
	r = l
	for r-l+1 < k {
		if l > 0 {
			if r+1 < len(arr) {
				if abs(arr[l-1], x) <= abs(arr[r+1], x) {
					l--
				} else {
					r++
				}
			} else {
				l--
			}
		} else {
			r++
		}
	}
	return arr[l : r+1]
}
