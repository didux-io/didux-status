package managers

import (
	"flag"
	"fmt"
	"math/big"
	"runtime"
	"smilo-status/models"
	"web3go/common"
	"web3go/provider"
	"web3go/rpc"
	"web3go/web3"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

var hostname = flag.String("hostname", "localhost", "The Smilo client RPC host")
//var hostname = flag.String("hostname", "18.202.153.27", "The Smilo client RPC host -- For testing!")
var port = flag.String("port", "22000", "The smilo client RPC port")
var verbose = flag.Bool("verbose", false, "Print verbose messages")
var defaultAccount common.Address
var connectedPeers uint64
var blockHeight *big.Int
var adminInfo *common.NodeInfo
var peers []common.Peer
var txpool *common.Txpool
var err error


//Return Status.
func GetStatus() models.Status {

	flag.Parse()

	if *verbose {
		fmt.Printf("Connect to %s:%s\n", *hostname, *port)
	}

	/**
	 * Connecting to provider with web3go
	 */
	provider := provider.NewHTTPProvider(*hostname+":"+*port, rpc.GetDefaultMethod())
	web3 := web3.NewWeb3(provider)

	// Get TxPool
	if txpool, err = web3.Txpool.Status(); err != nil {
		fmt.Printf("%v", err)
	}

	// Get Coinbase
	if defaultAccount, err = web3.Eth.Coinbase(); err != nil {
		fmt.Printf("%v", err)
	}

	// Get blockheight
	if blockHeight, err = web3.Eth.BlockNumber(); err != nil {
		fmt.Printf("%v", err)
	}

	// Get nodeInfo
	if adminInfo, err = web3.Admin.NodeInfo(); err != nil {
		fmt.Printf("%v", err)
	}

	// Get peers
	if peers, err = web3.Admin.Peers(); err != nil {
		fmt.Printf("Error: %v", err)
	}

	// Get connected peers
	if connectedPeers, err = web3.Net.PeerCount(); err != nil {
		fmt.Printf("%v", err)
	}

	// Return status of localhost here!!!
	return models.Status{Network: "Smilo", Address: defaultAccount.String(), BlockHeight: blockHeight, PeerCount: connectedPeers, Txpool: txpool, NodeInfo: adminInfo, ConnectedPeers: peers, System: GetSystemStatus()}
}

//Return Status.
func GetSystemStatus() models.System {
	runtimeOS := runtime.GOOS

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("%v", err)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		fmt.Printf("%v", err)
	}

	hostStat, err := host.Info()
	if err != nil {
		fmt.Printf("%v", err)
	}

	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	if err != nil {
		fmt.Printf("%v", interfStat)
	}

	return models.System{OS: runtimeOS, Host: hostStat.Hostname, Uptime: hostStat.Uptime, Processes: hostStat.Procs, TotalMemory: vmStat.Total, FreeMemory: vmStat.Free, MemoryUsage: vmStat.UsedPercent, TotalDiskSpace: diskStat.Total, UsedDiskSpace: diskStat.Used, FreeDiskSpace: diskStat.Free, DiskSpaceUsage: diskStat.UsedPercent}
}