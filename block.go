package main

import (
	"encoding/json"
	"time"
)

type Header struct {
	PrevHash int
	Timestamp int64
	Difficulty uint
	Nonce uint
	// version
	// merkleRoot
}

type Block struct {
	// size uint
	Header *Header
	Hash string
	Data interface{}
}

func newBlock() *Block {
	return &Block {
		Header: &Header{
			PrevHash: nil,
			Timestamp: time.Now().Unix(),
			Difficulty: nil,
			Nonce: nil,
		},
		Hash: nil,
		Data: []byte(""),
	}
}
