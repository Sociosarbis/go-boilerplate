package special

import "sort"

func makeLargestSpecial(s string) string {
	mounts := []string{}

	for i := 0; i < len(s); i++ {
		count := 0
		j := i
		for ; j < len(s); j++ {
			if s[j] == '1' {
				count++
			} else {
				count--
			}
			if count == 0 {
				// 这里的递归能够保证即是连续的又是特殊的，特殊字符串必然是1开头，0结尾，这样就能把问题缩小，进行递归
				mounts = append(mounts, "1"+makeLargestSpecial(s[i+1:j])+"0")
				i = j + 1
			}
		}
	}
	sort.Strings(mounts)
	ret := ""
	for i := len(mounts) - 1; i >= 0; i-- {
		ret += mounts[i]
	}
	return ret
}
