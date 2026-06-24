package arrays2

const mask int = 1e9 + 7

func matMul(a [][]int, b [][]int, out [][]int) {
	m := len(out)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			var temp int
			for k := 0; k < m; k++ {
				temp = (temp + int(int64(a[i][k])*int64(b[k][j])%int64(mask))) % mask
			}
			out[i][j] = temp
		}
	}
}

func qpow(mat [][]int, k int) [][]int {
	m := len(mat)
	out, temp := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		out[i], temp[i] = make([]int, m), make([]int, m)
		out[i][i] = 1
	}
	for k != 0 {
		if k&1 == 1 {
			for i := 0; i < m; i++ {
				copy(temp[i], out[i])
			}
			matMul(temp, mat, out)
			k--
		} else {
			matMul(mat, mat, temp)
			for i := 0; i < m; i++ {
				copy(mat[i], temp[i])
			}
			k >>= 1
		}
	}
	return out
}

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1
	mm := 2 * m
	mat := make([][]int, mm)
	for i := 0; i < mm; i++ {
		mat[i] = make([]int, mm)
	}
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			mat[i][m+j] = 1
		}
		for j := 0; j < i; j++ {
			mat[i+m][j] = 1
		}
	}
	mat = qpow(mat, n-1)
	var out int
	for i := 0; i < mm; i++ {
		for j := 0; j < mm; j++ {
			out = (out + mat[i][j]) % mask
		}
	}
	return out
}
