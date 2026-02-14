package balanced4

import "fmt"

const base int64 = 100003

func hash(counter []int) int64 {
	min := getMin2(counter)
	var h int64
	for _, c := range counter {
		h = h*base + int64(c-min)
	}
	return h
}

func hash2(counter []int) int64 {
	var h int64
	for _, c := range counter {
		h = h*base + int64(c)
	}
	return h
}

func isBalanced(counter []int) bool {
	min := getMin(counter)
	for _, c := range counter {
		if c > 0 {
			if min != c {
				return false
			}
		}
	}
	return true
}

// 非0的最小值
func getMin(counter []int) int {
	var min int
	for _, c := range counter {
		if c > 0 {
			if min == 0 || c < min {
				min = c
			}
		}
	}
	return min
}

// 最小值
func getMin2(counter []int) int {
	min := counter[0]
	for i := 1; i < 3; i++ {
		if counter[i] < min {
			min = counter[i]
		}
	}
	return min
}

func getCount(counter []int) int {
	var count int
	for _, c := range counter {
		if c > 0 {
			count++
		}
	}
	return count
}

func hash3(counter []int) int64 {
	min := getMin(counter)
	c2 := make([]int, 3)
	for i, c := range counter {
		if c > min {
			c2[i] = c - min
		}
	}
	return hash2(c2)
}

func longestBalanced(s string) int {
	n := len(s)
	m3 := make(map[int64]int, n)
	m2 := make(map[int64]map[int64]int, n)
	m1 := make(map[int64]int, n)
	counter := make([]int, 3)
	var out int
	for i, c := range s {
		counter[c-'a']++
		index := i
		if isBalanced(counter) {
			out = i + 1
		} else {
			// 保留1个
			for i := 0; i < 3; i++ {
				c1 := make([]int, 3)
				for j := 0; j < 3; j++ {
					if j != i {
						c1[j] = counter[j]
					}
				}
				h1 := hash2(c1)
				if j, ok := m1[h1]; ok {
					if index-j > out {
						fmt.Println("1", index-j)
						out = index - j
					}
				}
			}
			// 保留2个
			for i := 0; i < 3; i++ {
				// 保留的数
				c2 := make([]int, 3)
				// 减掉的数
				c21 := make([]int, 3)
				for j := 0; j < 3; j++ {
					if j != i {
						c2[j] = counter[j]
					} else {
						c21[j] = counter[j]
					}
				}

				h21 := hash2(c21)
				h2 := hash3(c2)
				if _, ok := m2[h21]; ok {
					if j, ok := m2[h21][h2]; ok {
						if index-j > out {
							fmt.Println("2", index-j)
							out = index - j
						}
					}
				}
			}
			// 保留3个
			h3 := hash3(counter)
			if j, ok := m3[h3]; ok {
				if index-j > out {
					fmt.Println("3", index-j)
					out = index - j
				}
			}
		}

		for i := 0; i < 3; i++ {
			// 减掉的数
			c1 := make([]int, 3)
			for j := 0; j < 3; j++ {
				if j != i {
					c1[j] = counter[j]
				}
			}
			h1 := hash2(c1)
			if _, ok := m1[h1]; !ok {
				m1[h1] = index
			}
		}
		for i := 0; i < 3; i++ {
			// 保留的数
			c2 := make([]int, 3)
			// 减掉的数
			c21 := make([]int, 3)
			for j := 0; j < 3; j++ {
				if j != i {
					c2[j] = counter[j]
				} else {
					c21[j] = counter[j]
				}
			}

			h21 := hash2(c21)
			if _, ok := m2[h21]; !ok {
				m2[h21] = map[int64]int{}
			}
			var h2 int64
			if getCount(c2) == 1 {
				h2 = hash2(c2)
			} else {
				h2 = hash3(c2)
			}
			if _, ok := m2[h21][h2]; !ok {
				m2[h21][h2] = index
			}
		}
		h3 := hash(counter)
		if _, ok := m3[h3]; !ok {
			m3[h3] = index
		}
		fmt.Println(i, out)
	}
	return out
}
