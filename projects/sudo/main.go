package main

import (
	"fmt"
)

var txt = `008090000
070000280
064100309
000805900
500000001
009304000
802007560
097000010
000060700`

func makeSudo(s string) sudo {
	var su sudo
	for i := range s {
		if (i + 1) % 10 != 0 {
			row := i / 10
			col := i % 10
			num := s[i]
			su[row][col] = initBlock(num, row, col, &su)
		}
	}
	return su
}

func in(b byte, line []byte) bool {
	for _, i := range line {
		if i == b {
			return true
		}
	}
	return false
}

func scan(su *sudo) int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !su[i][j].IsOk() {
				if len(su[i][j].Maybe()) == 1 {
					su[i][j].num = su[i][j].Maybe()[0]
					if su.IsOk() {
						return 0
					}
					scan(su)
				} else if len(su[i][j].Maybe()) == 0 {
					return -1
				}
			}
		}
	}
	return 1
}

func initNode() node {
	var top node
	var node2 node
	node2.pre = &top
	node2.floor = 0
	node2.offset = 0
	return node2
}


func main() {
	sudo001 := makeSudo(txt)
	maylist2 := sudo001.MakeMayList()
	temp := maylist2.Gen()
	for {
		result := temp(1)
		fmt.Println(result)
	}
}
