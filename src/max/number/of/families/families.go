package families

import "sort"

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	sort.Slice(reservedSeats, func(i, j int) bool {
		r1, c1, r2, c2 := reservedSeats[i][0], reservedSeats[i][1], reservedSeats[j][0], reservedSeats[j][1]
		return r1 < r2 || (r1 == r2 && c1 < c2)
	})
	cr, l := 1, 1
	var out int
	for _, seat := range reservedSeats {
		r, c := seat[0], seat[1]
		if cr != r {
			if l <= 6 {
				if l > 2 {
					out++
				} else {
					out += 2
				}
			}
			l = 1
			out += (r - 1 - cr) * 2
			cr = r
		}
		if l <= 6 && c >= l+3 {
			if l > 4 {
				if c == 10 {
					out++
				}
			} else if l > 2 {
				if c > 7 {
					out++
				}
			} else {
				if c == 10 {
					out += 2
				} else if c > 5 {
					out++
				}
			}
		}
		l = c + 1
	}
	if cr <= n {
		if l <= 6 {
			if l > 2 {
				out++
			} else {
				out += 2
			}
		}
		out += 2 * (n - cr)
	}
	return out
}
