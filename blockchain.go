package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

const (
	MiningDifficulty = 4
)

type Blockchain struct {
	Blocks []*Block
}

func newBlockChain() *Blockchain {
	// create and mine genesis block
	prevHash := "0000000000000000000000000000000000000000000000000000000000000000"
	blockData := "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."
	proof, nonce := mine(blockData)
	genesisBlock := newBlock(prevHash, nonce, blockData, proof)

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func mine(data string) (string, int) {
	nonce := 0
	for {
		// convert nonce and data to byte arrays, append, and compute hashes
		h := sha256.New()
		nonceBytes := []byte(strconv.Itoa(nonce))
		h.Write(append([]byte(data), nonceBytes...))
		hash := hex.EncodeToString(h.Sum(nil))

		var difficulty bytes.Buffer
		for i := 0; i < MiningDifficulty; i++ {
			difficulty.WriteString("0")
		}
		// check if valid hash (ie number of zeroes at beginning of hash equals MiningDifficulty)
		if hash[0:MiningDifficulty] == difficulty.String() {
			return hash, nonce
		}

		nonce = nonce + 1
	}
}
