package main

type block struct {
	num   byte
	row   int
	col   int
	owner *sudo
}

func initBlock(num byte, row, col int, su *sudo) block {
	var b block
	b.num = num
	b.row = row
	b.col = col
	b.owner = su
	return b
}

func (b block) IsOk() bool {
	return b.num != '0'
}

func (b block) GetRow() []byte {
	var row []byte
	for i := 0; i < 9; i++ {
		row = append(row, b.owner[b.row][i].num)
	}
	return row
}

func (b block) GetCol() []byte {
	var col []byte
	for i := 0; i < 9; i++ {
		col = append(col, b.owner[i][b.col].num)
	}
	return col
}

func (b block) GetTup() []byte {
	var result []byte
	baseRow := b.row / 3 * 3
	baseCol := b.col / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			result = append(result, b.owner[baseRow+i][baseCol+j].num)
		}
	}
	return result
}

func (b block) Maybe() []byte {
	if b.IsOk() {
		return []byte{b.num}
	}
	var result []byte
	for i := byte('1'); i <= '9'; i++ {
		if !in(i, b.GetRow()) && !in(i, b.GetCol()) && !in(i, b.GetTup()) {
			result = append(result, i)
		}
	}
	return result
}
