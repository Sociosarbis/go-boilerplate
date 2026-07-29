package palindrome2

func comb(n, k, bound int) int {
	out := 1
	// 组合数 c(n, k) == c(n, n - k)
	// 这样可以让下面的循环，out是单调递增的
	if k+k > n {
		k = n - k
	}
	for i := 0; i < k; i++ {
		// 连续n个数，一定能够被i!整除
		out = out * (n - i) / (i + 1)
		if out > bound {
			return out
		}
	}
	return out
}

func smallestPalindrome(s string, k int) string {
	n := len(s) / 2
	counter := [26]int{}
	for _, c := range s {
		counter[c-'a']++
	}
	oddIndex := -1
	for i := 0; i < 26; i++ {
		if counter[i]&1 != 0 {
			oddIndex = i
			break
		}
	}
	for i := 0; i < 26; i++ {
		if counter[i] != 0 {
			counter[i] >>= 1
		}
	}
	out := make([]byte, 0, len(s))
block:
	for i := 1; i < n; i++ {

		for j := 0; j < 26; j++ {
			if counter[j] != 0 {
				counter[j]--
				temp := 1
				tempN := n - i
				for l := 0; l < 26; l++ {
					if counter[l] != 0 {
						// 从remain中取counter[l]个位置，放置l，等于C(remain, counter[l])
						temp *= comb(tempN, counter[l], k)
						if temp >= k {
							out = append(out, 'a'+byte(j))
							continue block
						}
						tempN -= counter[l]
					}
				}
				counter[j]++
				k -= temp
			}
		}
		return ""
	}
	if k > 1 {
		return ""
	}
	for i := 0; i < 26; i++ {
		if counter[i] != 0 {
			out = append(out, 'a'+byte(i))
		}
	}
	if oddIndex != -1 {
		out = append(out, 'a'+byte(oddIndex))
	}
	for i := n - 1; i >= 0; i-- {
		out = append(out, out[i])
	}
	return string(out)
}
