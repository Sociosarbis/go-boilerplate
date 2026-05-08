package jumps

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func minJumps(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	prime := make(map[int][]int, n)
	for _, num := range nums {
		if _, ok := prime[num]; ok {
			continue
		}
		if isPrime(num) {
			prime[num] = []int{}
		}
	}
	for i, num := range nums {
		if num > 1 {
			if next, ok := prime[num]; ok {
				prime[num] = append(next, i)
			} else {
				for j := 2; j*j <= num; j++ {
					if num%j == 0 {
						if next, ok := prime[j]; ok {
							prime[j] = append(next, i)
						}
						nj := num / j
						if nj != j {
							if next, ok := prime[nj]; ok {
								prime[nj] = append(next, i)
							}
						}
					}
				}
			}
		}
	}
	bfs := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	var out int
	m := 1
	for m > 0 {
		out++
		for i := 0; i < m; i++ {
			index := bfs[i]
			if next, ok := prime[nums[index]]; ok {
				for _, j := range next {
					if !visited[j] {
						bfs = append(bfs, j)
						visited[j] = true
					}
				}
				delete(prime, nums[index])
			}
			for j := -1; j < 2; j += 2 {
				if index+j >= 0 && index+j < n && !visited[index+j] {
					bfs = append(bfs, index+j)
					visited[index+j] = true
				}
			}
		}
		if visited[n-1] {
			return out
		}
		bfs = bfs[m:]
		m = len(bfs)
	}
	return out
}
