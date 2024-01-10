package main

type Chain struct {
	Blocks []*Block
}

func (c *Chain) AddBlock(data string) {
	prevBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}

func NewChain() *Chain {
	return &Chain{[]*Block{AddGenesisBlock()}}
}
