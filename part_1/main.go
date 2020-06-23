package main

import (
	"fmt"

)




func main() {
	bc := NewBlockchain();
	bc.AddBlock("send BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	
	for _, block := range bc.blocks {
		fmt.Printf("Prev . hash %x\n", block.PrevBlockHach)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		fmt.Println()
	}
}
