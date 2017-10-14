package main

import (
	"time"
)

type Header struct {
	PrevHash   string
	Timestamp  int64
	Difficulty uint
	Nonce      uint
}

type Block struct {
	Header *Header
	Hash   string
	Data   interface{}
}

func newBlock(prevHash string, nonce uint, data interface{}, proof string) *Block {
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
