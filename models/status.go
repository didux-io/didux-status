package models

import (
	"math/big"
)

type Status struct {
	Network			string
	NetworkId		*big.Int
	Address   		string
	BlockHeight 	*big.Int
	PeerCount     	uint64
	PendingTransactions uint
}

type Result struct {
	Message string
}
