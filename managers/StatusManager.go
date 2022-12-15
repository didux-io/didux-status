package managers

import (
	"didux-status/models"
	"didux-status/web3go/common"
	"didux-status/web3go/provider"
	"didux-status/web3go/rpc"
	"didux-status/web3go/web3"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var hostname = flag.String("hostname", "37.59.131.19", "The Didux client RPC host")

// var hostname = flag.String("hostname", "18.202.153.27", "The Didux client RPC host -- For testing!")
var port = flag.String("port", "22000", "The didux client RPC port")
var verbose = flag.Bool("verbose", false, "Print verbose messages")
var defaultAccount common.Address
var connectedPeers uint64
var blockHeight *big.Int
var adminInfo *common.NodeInfo
var peers []common.Peer
var txpool *common.Txpool
var err error
var statusCode = http.StatusNoContent
var localBlockStatus int

// Return Status.
func GetStatus() models.Status {
	return models.Status{Didux: GetDidux(), System: GetSystemStatus()}
}

func GetStatusCode() int {
	return statusCode
}

// Return Status.
func GetBlock() models.Block {
	return models.Block{BlockHeight: GetBlockHeight()}
}

func GetPublicBlocks() {
	localBlockHeight := GetBlockHeight()
	publicNodes := [4]string{"http://212.32.245.83:5000/block", "http://212.32.245.91:5000/block", "http://51.89.103.235:5000/block", "http://37.59.131.19:5000/block"}
	localBlockStatus = http.StatusOK

	if localBlockHeight == nil {
		fmt.Println("Could not fetch localBlocks")
		localBlockStatus = http.StatusConflict
	} else {
		for _, url := range publicNodes {

			nodeClient := http.Client{
				Timeout: time.Second * 2, // Timeout after 2 seconds
			}

			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Set("User-Agent", "status-watcher")

			res, getErr := nodeClient.Do(req)
			if getErr != nil {
				log.Fatal(getErr)
			}

			if res.Body != nil {
				defer res.Body.Close()
			}

			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}

			blockHeight := models.Block{}
			jsonErr := json.Unmarshal(body, &blockHeight)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}

			// -1 = Node is behind publicNodes
			// 0 = Node is in sync with publicNodes
			// 1 = Node is ahead of publicNodes
			// Add 2 blocks to localhost, so we have a delta of +2, then compare with remote.
			// We expect the reomte server to be equal or lower than us.
			if localBlockHeight.Add(localBlockHeight, big.NewInt(2)).Cmp(blockHeight.BlockHeight) < 0 {
				fmt.Printf("Localhost is behind!\n")
				localBlockStatus = http.StatusConflict
			}
		}
	}
	statusCode = localBlockStatus
}

func GetBlockHeight() *big.Int {
	flag.Parse()

	if *verbose {
		fmt.Printf("Connect to %s:%s\n", *hostname, *port)
	}

	/**
	 * Connecting to provider with web3go
	 */
	provider := provider.NewHTTPProvider(*hostname+":"+*port, rpc.GetDefaultMethod())
	web3 := web3.NewWeb3(provider)

	// Get blockheight
	if blockHeight, err = web3.Eth.BlockNumber(); err != nil {
		fmt.Printf("%v", err)
	}

	return blockHeight
}

// Return Go-Didux overview
func GetDidux() models.Didux {

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
	return models.Didux{Network: "Didux", Address: defaultAccount.String(), BlockHeight: blockHeight, PeerCount: connectedPeers, Txpool: txpool, NodeInfo: adminInfo, ConnectedPeers: peers}
}

// Return Status.
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

	loadStats, err := load.Avg()
	if err != nil {
		fmt.Printf("%v", err)
	}

	hostInfo, err := host.Info()

	return models.System{OS: runtimeOS, Host: hostInfo.Hostname, Uptime: hostInfo.Uptime,
		Load1m: loadStats.Load1, Load5m: loadStats.Load5, Load15m: loadStats.Load15, Processes: hostInfo.Procs,
		TotalMemory: vmStat.Total, FreeMemory: vmStat.Free, MemoryUsage: vmStat.UsedPercent,
		TotalDiskSpace: diskStat.Total, UsedDiskSpace: diskStat.Used, FreeDiskSpace: diskStat.Free, DiskSpaceUsage: diskStat.UsedPercent}
}
