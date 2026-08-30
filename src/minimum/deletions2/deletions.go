package deletions2

func minimumDeletions(nums []int) int {
	n := len(nums)
	var min int = 1e5 + 1
	var max int = -1e5 - 1
	var mni, mxi int
	for i, num := range nums {
		if num < min {
			min = num
			mni = i
		}
		if num > max {
			max = num
			mxi = i
		}
	}

	if mni == mxi {
		if mni+1 < n-mxi {
			return mni + 1
		}
		return n - mxi
	}
	if mni > mxi {
		mni, mxi = mxi, mni
	}
	out := mni + 1 + n - mxi
	if n-mni < out {
		out = n - mni
	}
	if mxi+1 < out {
		return mxi + 1
	}
	return out
}
