package matrix

const mask int64 = 1e9 + 7

type mat struct {
	data [][]int64
}

func NewMat(rows int, cols int) mat {
	data := make([][]int64, rows)
	for i := 0; i < cols; i++ {
		data[i] = make([]int64, cols)
	}
	return mat{
		data,
	}
}

func NewIdentMat(rows int, cols int) mat {
	mat := NewMat(rows, cols)
	for i := 0; i < rows; i++ {
		mat.data[i][i] = 1
	}
	return mat
}

func (self *mat) Get(i, j int) int {
	return int(self.data[i][j])
}

func (self *mat) Set(i, j, value int) {
	self.data[i][j] = int64(value)
}

func (self *mat) Mul(other *mat) mat {
	m1 := len(self.data)
	m2 := len(other.data)
	n2 := len(other.data[0])
	out := NewMat(m1, n2)
	for i := 0; i < m1; i++ {
		for j := 0; j < n2; j++ {
			for k := 0; k < m2; k++ {
				out.data[i][j] = (out.data[i][j] + self.data[i][k]*other.data[k][j]) % mask
			}
		}
	}
	return out
}
