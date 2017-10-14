package main

import (
//"crypto/sha256"
)

const (
	MiningDifficulty = 2
)

type Blockchain struct {
	Blocks []*Block
}

func newBlockChain() *Blockchain {
	// create genesis block
	prevHash := "0000000000000000000000000000000000000000000000000000000000000000"
	blockData := "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."
	proof, nonce = mine(blockData)
	genesisBlock := newBlock(prevHash, nonce, blockData, proof)

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func mine(data interface{}) (string, uint) {

}

func isValidProof(proof string) bool {

}
