package main

import "fmt"

type cat struct {
	su sudo
	seq int
}

func (c *cat)Call() sudo {
	newSu := c.su.Copy()
	newSu.Locate().num = newSu.Locate().Maybe()[c.seq]
	return newSu
}

type stack []cat

func (s *stack)Pop() {
	*s = (*s)[:len(*s) -1]
}

func (s *stack)Push(c cat) {
	*s = append((*s), c)
}

func (s *stack)Last() (*cat) {
	fmt.Print(*s.Last())
	return &(*s)[len(*s)-1]
}

func (s *stack)deepScan() bool {
	for {
		if len(*s) == 0 {
			return false
		}
		if (*s).Last().seq == -1 {
			result := scan(&(s.Last().su))
			if result == 1 {
				(s.Last().su).Reader()
				return true
			} else if result == -1 {
				s.Pop()
				s.Last().seq++
			} else {
				s.Last().seq++
				s.Push(cat{su:s.Last().su.Copy(), seq:-1})
			}
		} else if ((*s).Last().seq) == len((*s).Last().su.Locate().Maybe()) {
			s.Pop()
		} else {
			check := s.Last().Call()
			result := scan(&check)
			if result == 1 {
				check.Reader()
				return true
			} else if result == -1 {
				s.Pop()
				s.Last().seq++
			} else {
				s.Last().seq++
				s.Push(cat{su:s.Last().su.Copy(), seq:-1})
			}
		}
	}
}