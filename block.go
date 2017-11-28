package main

type Block struct {
	Timestamp         int64
	PreviousBlockHash []byte
	Hash              []byte
}
