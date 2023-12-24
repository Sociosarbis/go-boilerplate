package perimeter

func minimumPerimeter(neededApples int64) int64 {
	var l int64 = 1
	var r int64 = 1e6
	for l <= r {
		mid := (l + r) / 2
		res := getRes(mid)
		if res < neededApples {
			l = mid + 1
		} else {
			if mid > 0 && getRes(mid-1) >= neededApples {
				r = mid - 1
			} else {
				l = mid
				break
			}
		}
	}
	return l * 8
}

func getRes(n int64) int64 {
	return (n + 1) * n * (2*n + 1) * 2
}
