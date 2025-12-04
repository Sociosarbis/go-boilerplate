package trapezoids2

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func normalize(x, y int) (int, int) {
	if x == 0 {
		y = 1
	} else if y == 0 {
		x = 1
	} else {
		if x < 0 {
			x, y = -x, -y
		}
		g := gcd(x, abs(y))
		x /= g
		y /= g
	}
	return x, y
}

func hash(x, y int) int {
	x, y = normalize(x, y)
	return x + 1000 + (y+1000)*2001
}

func countTrapezoids(points [][]int) int {
	n := len(points)
	buckets := make(map[int][][2]int, n*n)
	for i, point := range points {
		for j := i + 1; j < n; j++ {
			h := hash(points[j][0]-point[0], points[j][1]-point[1])
			if bucket, ok := buckets[h]; ok {
				buckets[h] = append(bucket, [2]int{i, j})
			} else {
				buckets[h] = [][2]int{{i, j}}
			}
		}
	}
	var out int64
	var subtract int64
	for _, bucket := range buckets {
		n = len(bucket)
		if n == 1 {
			continue
		}
		temp := int64(n) * int64(n-1) / 2
		ks := make(map[int]int, n)
		ls := make(map[int]map[int]int, n)
		pair := bucket[0]
		x, y := normalize(points[pair[1]][0]-points[pair[0]][0], points[pair[1]][1]-points[pair[0]][1])
		for _, pair := range bucket {
			k := x*points[pair[0]][1] - y*points[pair[0]][0]
			if c, ok := ks[k]; ok {
				ks[k] = c + 1
			} else {
				ks[k] = 1
			}
			dx, dy := points[pair[1]][0]-points[pair[0]][0], points[pair[1]][1]-points[pair[0]][1]
			l := dx*dx + dy*dy
			if m, ok := ls[l]; ok {
				if c, ok := m[k]; ok {
					m[k] = c + 1
				} else {
					m[k] = 1
				}
				ls[l] = m
			} else {
				ls[l] = map[int]int{
					k: 1,
				}
			}
		}
		for _, v := range ks {
			if v > 1 {
				temp -= int64(v) * int64(v-1) / 2
			}
		}
		for _, v := range ls {
			var sum int
			for _, c := range v {
				sum += c
			}
			for _, c := range v {
				subtract += int64(sum-c) * int64(c)
				sum -= c
			}
		}
		out += temp
	}
	return int(out - subtract/2)
}
