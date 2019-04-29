package models

import (
	"math/big"
	"web3go/common"
)

type Status struct {
	Smilo 		Smilo
	System 		System
}

type Smilo struct {
	Network				string
	Address   			string
	BlockHeight 		*big.Int
	PeerCount     		uint64
	Txpool				*common.Txpool
	NodeInfo			*common.NodeInfo
	ConnectedPeers		[]common.Peer

}

type System struct {
	OS							string
	Host						string
	Uptime						uint64
	Load1m						float64
	Load5m						float64
	Load15m						float64
	Processes					uint64
	TotalMemory					uint64
	FreeMemory					uint64
	MemoryUsage					float64
	TotalDiskSpace				uint64
	UsedDiskSpace   			uint64
	FreeDiskSpace 				uint64
	DiskSpaceUsage	     		float64
}

type Result struct {
	Message string
}
