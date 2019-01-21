package main

import (
	"fmt"
	"errors"
)


type sudo [9][9]block

func (su sudo) Reader() {
	for i := 0; i < 9; i++ {
		var line []string
		for j := 0; j < 9; j++ {
			line = append(line, string(su[i][j].num))
		}
		fmt.Println(line)
	}
}

func (su *sudo) IsOk() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !su[i][j].IsOk() {
				return false
			}
		}
	}
	return true
}

func (su *sudo) Gen() func() (block, error) {
	r, c := 0, 0
	for {
		return func() (block, error) {
			if c == 9 {
				return block{num:47}, errors.New("OVER")
			}
			fmt.Println(r, c)
			result := su[r][c]
			if c == 8 && r < 8 {
				r += 1
				c = 0
			}
			c += 1
			return result, nil
		}
	}
}

func (su *sudo) Catch() func() block {
	temp := su.Gen()
	return func() block {
		for i := 2; i < 10; i++ {
			for {
				cat, _ := temp()
				if cat.num == 47 {
					break
				} else {
					if len(cat.Maybe()) == i {return cat}
				}
			}
		}
		return block{num:47}
	}
}
