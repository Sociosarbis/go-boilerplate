package swap

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	swap := make([]int, n)
	noSwap := make([]int, n)
	for i := 0; i < n; i++ {
		swap[i] = n
	}
	copy(noSwap, swap)
	swap[0] = 1
	noSwap[0] = 0
	// swap[i]表示使得0->i单调递增且交换位置i情况下的最小交换次数，noSwap[i]表示使得0->i单调递增且不交换位置i情况下的最小交换次数
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			swap[i] = swap[i-1] + 1
			noSwap[i] = noSwap[i-1]
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			if noSwap[i-1]+1 < swap[i] {
				swap[i] = noSwap[i-1] + 1
			}
			if swap[i-1] < noSwap[i] {
				noSwap[i] = swap[i-1]
			}
		}
	}
	if swap[n-1] < noSwap[n-1] {
		return swap[n-1]
	}
	return noSwap[n-1]
}
