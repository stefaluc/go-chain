package main

import (
	"fmt"
)

func main() {
	blockchain := newBlockChain()
	fmt.Printf("Hash: %s, Nonce: %d", blockchain.Blocks[0].Hash, blockchain.Blocks[0].Header.Nonce)
}
