package main

import (
	"fmt"
)


type sudo [9][9]block

func (su sudo)Reader() {
	for i := 0; i < 9; i++ {
		var line []string
		for j := 0; j < 9; j++ {
			line = append(line, string(su[i][j].num))
		}
		fmt.Println(line)
	}
}

func (su *sudo)IsOk() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !su[i][j].IsOk() {
				return false
			}
		}
	}
	return true
}

func (su *sudo)Locate() *block {
	for i:=2; i<10; i++ {
		for r:=0; r<9; r++ {
			for c:=0; c<9; c++{
				if len((*su)[r][c].Maybe()) == i{
					return &(*su)[r][c]
				}
			}
		}
	}
	return &block{row:-1, col:-1}
}

func (su *sudo)Copy() sudo {
	var newSu sudo
	for r:=0; r<9; r++{
		for c:=0; c<9; c++{
			var blk block
			blk.num = (*su)[r][c].num
			blk.row = r
			blk.col = c
			blk.owner = &newSu
			newSu[r][c] = blk
		}
	}
	return newSu
}

func (su *sudo)MakeMayList() maylist {
	var maylist2 maylist
	newSu := su.Copy()
	scan(&newSu)
	for length := 2; length<9; length++ {
		for r:=0; r<9; r++{
			for c:=0; c<9; c++{
				if len(newSu[r][c].Maybe()) == length {
					maylist2 = append(maylist2, may{row:r, col:c, may:newSu[r][c].Maybe()})
				}
			}
		}
	}
	return maylist2
}

func (su *sudo)Join(node2 node) sudo {
	newSu := su.Copy()
	maylist2 := su.MakeMayList()
	link := node2.Gen(maylist2)
	for floor, num := range link {
		newSu[maylist2[floor].row][maylist2[floor].col].num = num
	}
	return newSu
}