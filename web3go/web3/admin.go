// Copyright (c) 2016, Alan Chen
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package web3

import (
	"didux-status/web3go/common"
	"encoding/json"
	"fmt"
)

// Admin ...
type Admin interface {
	NodeInfo() (*common.NodeInfo, error)
	Peers() ([]common.Peer, error)
}

// AdminAPI ...
type AdminAPI struct {
	requestManager *requestManager
}

// NewNetAPI ...
func newAdminAPI(requestManager *requestManager) Admin {
	return &AdminAPI{requestManager: requestManager}
}

// NodeInfo returns the nodeInfo.
func (admin *AdminAPI) NodeInfo() (*common.NodeInfo, error) {
	req := admin.requestManager.newRequest("admin_nodeInfo")
	resp, err := admin.requestManager.send(req)
	if err != nil {
		return nil, err
	}
	result := &jsonNodeInfo{}
	if jsonBytes, err := json.Marshal(resp.Get("result")); err == nil {
		//fmt.Printf("%s", jsonBytes)
		if err := json.Unmarshal(jsonBytes, result); err == nil {
			return result.ToNodeInfo(), nil
		}
	}
	return nil, fmt.Errorf("%v", resp.Get("result"))
}

// Returns the connected peers.
func (admin *AdminAPI) Peers() ([]common.Peer, error) {
	req := admin.requestManager.newRequest("admin_peers")
	resp, err := admin.requestManager.send(req)
	if err != nil {
		return nil, err
	}

	jsonResponse, err := json.Marshal(resp.Get("result"))

	peerList := make([]common.Peer, 0)
	err = json.Unmarshal(jsonResponse, &peerList)
	if err == nil {
		//fmt.Printf("%+v\n", peerList)
		return peerList, nil
	} else {
		fmt.Println(err)
		fmt.Printf("%+v\n", peerList)
	}

	return nil, fmt.Errorf("%v", resp.Get("result"))
}
