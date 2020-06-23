package main

type Blocchain struct {
	blocks []*Block
}

func (bc *Blocchain) AddBlock(data string) {
	PrevBlockHach := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, PrevBlockHach.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func Newgenesisblock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blocchain {
	return &Blocchain{[]*Block{Newgenesisblock()}}
}
