package managers

import (
	"context"
	"flag"
	"fmt"
	"github.com/alanchchen/web3go/common"
	"github.com/alanchchen/web3go/provider"
	"github.com/alanchchen/web3go/rpc"
	"github.com/alanchchen/web3go/web3"
	"go-smilo/src/blockchain/smilobft/ethclient"
	"log"
	"math/big"
	"smilo-status/models"
)

var hostname = flag.String("hostname", "localhost", "The Smilo client RPC host")
//var hostname = flag.String("hostname", "18.202.153.27", "The Smilo client RPC host")
var port = flag.String("port", "22000", "The smilo client RPC port")
var verbose = flag.Bool("verbose", false, "Print verbose messages")
var defaultAccount common.Address
var connectedPeers uint64
var blockHeight *big.Int
var version string


//Return Status.
func GetStatus() models.Status {

	flag.Parse()

	if *verbose {
		fmt.Printf("Connect to %s:%s\n", *hostname, *port)
	}

	/**
	 * Connecting to provider with ethclient
	 */
	ctx := context.Background()
	client, err := ethclient.Dial("http://"+*hostname+":"+*port)

	if err != nil {
		log.Fatal(err)
	}

	networkId, err := client.NetworkID(ctx)
	pendingTransactions, err := client.PendingTransactionCount(ctx)

	/**
	 * Connecting to provider with web3go
	 */
	provider := provider.NewHTTPProvider(*hostname+":"+*port, rpc.GetDefaultMethod())
	web3 := web3.NewWeb3(provider)

	// Get default account
	if accounts, err := web3.Eth.Accounts(); err == nil {
		for _, account := range accounts {
			if defaultAccount.String() == "0x0000000000000000000000000000000000000000" {
				defaultAccount = account
			}
		}
	} else {
		fmt.Printf("%v", err)
	}

	// Get blockheight
	if blockHeight, err = web3.Eth.BlockNumber(); err != nil {
		fmt.Printf("%v", err)
	}

	// Get connected peers
	if connectedPeers, err = web3.Net.PeerCount(); err != nil {
		fmt.Printf("%v", err)
	}

	// Return status of localhost here!!!
	return models.Status{Network: "Smilo", NetworkId: networkId, Address: defaultAccount.String(), BlockHeight: blockHeight, PeerCount: connectedPeers, PendingTransactions: pendingTransactions }
}
