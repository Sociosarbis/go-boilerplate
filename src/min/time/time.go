package time

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	f := make([]int64, n)
	for i := 0; i < m; i++ {
		now := f[0]
		// 先得到满足所有最早时间条件的完成时间，然后反推第ith瓶药水至少得从什么时间开始制作
		for j := 1; j < n; j++ {
			temp := now + int64(skill[j-1])*int64(mana[i])
			if temp > f[j] {
				now = temp
			} else {
				now = f[j]
			}
		}
		f[n-1] = now + int64(skill[n-1])*int64(mana[i])
		for j := n - 2; j >= 0; j-- {
			f[j] = f[j+1] - int64(skill[j+1])*int64(mana[i])
		}
	}
	return f[n-1]
}
