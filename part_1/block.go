package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHach []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	Timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHach, b.Data, Timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, PrevBlockHach []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), PrevBlockHach, []byte{}}
	block.SetHash()
	return block
}
