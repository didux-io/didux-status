package models

import (
	"web3go/common"
	"math/big"
)

type Status struct {
	Network				string
	Address   			string
	BlockHeight 		*big.Int
	PeerCount     		uint64
	Txpool				*common.Txpool
	NodeInfo			*common.NodeInfo
	ConnectedPeers		[]common.Peer
}

type Result struct {
	Message string
}
