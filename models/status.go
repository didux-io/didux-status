package models

import (
	"web3go/common"
	"math/big"
)

type Status struct {
	Network				string
	NetworkId			*big.Int
	Address   			string
	BlockHeight 		*big.Int
	PeerCount     		uint64
	PendingTransactions uint
	NodeInfo			*common.NodeInfo
	ConnectedPeers		[]common.Peer
}

type Result struct {
	Message string
}
