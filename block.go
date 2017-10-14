package main

import (
	"time"
)

type Header struct {
	PrevHash   string
	Timestamp  int64
	Difficulty uint
	Nonce      int
}

type Block struct {
	Header *Header
	Hash   string
	Data   string
}

func newBlock(prevHash string, nonce int, data string, proof string) *Block {
	return &Block{
		Header: &Header{
			PrevHash:   prevHash,
			Timestamp:  time.Now().Unix(),
			Difficulty: MiningDifficulty,
			Nonce:      nonce,
		},
		Hash: proof,
		Data: data,
	}
}
