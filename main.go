package main

import (
	"fmt"

	"github.com/hellicopthecat/learngo/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second-Blcok")
	chain.AddBlock("Third-Blcok")
	chain.AddBlock("Fourth-Blcok")
	for _, block := range chain.AllBlocks() {
		fmt.Println(block.Data)
	}
}
