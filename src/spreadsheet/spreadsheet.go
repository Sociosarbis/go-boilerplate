package spreadsheet

type Spreadsheet struct {
	grid [][26]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		grid: make([][26]int, rows),
	}
}

func parseInt(value string) int {
	var out int
	for _, c := range value {
		out = out*10 + int(c-48)
	}
	return out
}

func getRowCol(cell string) (int, int) {
	return parseInt(cell[1:]) - 1, int(cell[0] - 'A')
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	row, col := getRowCol(cell)
	this.grid[row][col] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	row, col := getRowCol(cell)
	this.grid[row][col] = 0
}

func (this *Spreadsheet) parseCellValueOrValue(s string) int {
	if s[0] >= 'A' && s[0] <= 'Z' {
		row, col := getRowCol(s)
		return this.grid[row][col]
	}
	return parseInt(s)
}

func (this *Spreadsheet) getCellValueOrValue(s string) (int, int) {
	var i int
	n := len(s)
	for ; i < n; i++ {
		if s[i] == '+' {
			break
		}
	}
	return this.parseCellValueOrValue(s[0:i]), i
}

func (this *Spreadsheet) GetValue(formula string) int {
	a, i := this.getCellValueOrValue(formula[1:])
	b := this.parseCellValueOrValue(formula[i+2:])
	return a + b
}
