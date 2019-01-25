package main

type may struct{
	row int
	col int
	may []byte
}

type maylist []may

type node struct {
	pre *node
	floor int
	offset int
}

func (baseNode node)Gen(maylist2 maylist) []byte {
	var bytes []byte
	var  preNode node
	for i:=baseNode.floor; i>=0; i-- {
	preNode = baseNode
	for j:=0; j<i; j++ {
		preNode = *(preNode.pre)
	}
		bytes = append(bytes, maylist2[preNode.floor].may[preNode.offset])
	}
	return bytes
}

func (baseNode node)Next(forward int, maylist2 maylist) []node {
	if forward == 1 && baseNode.offset == len(maylist2[baseNode.floor].may) {
		nextNode := baseNode
		nextNode.offset++
		return nextNode.Next(1, maylist2)
	}
	nextNode := *baseNode.pre
	return nextNode.Next(1, maylist2)
}