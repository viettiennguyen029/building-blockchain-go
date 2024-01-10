package main

import (
	"fmt"
	"strconv"
)

func main() {
	testChain := NewChain()

	testChain.AddBlock("First Block")
	testChain.AddBlock("Second Block")

	for _, block := range testChain.Blocks {
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
